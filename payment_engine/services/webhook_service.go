package payment_engine_service

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	contracts "sharpnr-engine/payment_engine/contracts"
	models "sharpnr-engine/payment_engine/models"

	"github.com/google/uuid"
)

type WebhookEvent struct {
	EventType string         `json:"event_type"`
	EventID   string         `json:"event_id"`
	Payload   map[string]any `json:"payload"`
}

func (s *paymentEngine) RegisterWebhookEndpoint(
	ctx context.Context,
	req contracts.RegisterWebhookRequest,
) (*models.WebhookEndpoint, error) {

	if req.URL == "" {
		return nil, fmt.Errorf("webhook URL is required")
	}

	eventsJSON := "[]"
	if len(req.SubscribedEvents) > 0 {
		raw, _ := json.Marshal(req.SubscribedEvents)
		eventsJSON = string(raw)
	}

	endpoint := &models.WebhookEndpoint{
		ID:               uuid.New(),
		TenantID:         req.TenantID,
		URL:              req.URL,
		SigningSecret:    generateSigningSecret(),
		SubscribedEvents: eventsJSON,
		IsActive:         true,
	}

	if err := s.paymentRepo.CreateWebhookEndpoint(ctx, endpoint); err != nil {
		return nil, fmt.Errorf("create webhook endpoint: %w", err)
	}
	return endpoint, nil
}

func (s *paymentEngine) ListWebhookEndpoints(
	ctx context.Context,
	tenantID uuid.UUID,
) ([]*models.WebhookEndpoint, error) {

	endpoints, err := s.paymentRepo.ListWebhookEndpoints(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("list webhook endpoints: %w", err)
	}
	for _, e := range endpoints {
		e.SigningSecret = "" // never leak
	}
	return endpoints, nil
}

func (s *paymentEngine) RemoveWebhookEndpoint(
	ctx context.Context,
	id uuid.UUID,
) error {

	if _, err := s.paymentRepo.GetWebhookEndpointByID(ctx, id); err != nil {
		return fmt.Errorf("webhook endpoint not found: %w", err)
	}
	return s.paymentRepo.DeleteWebhookEndpoint(ctx, id)
}

// DispatchEvent — fire and forget, never blocks the payment flow
func (s *paymentEngine) DispatchEvent(
	ctx context.Context,
	tenantID uuid.UUID,
	event WebhookEvent,
) {
	endpoints, err := s.paymentRepo.GetActiveWebhookEndpoints(ctx, tenantID, event.EventType)
	if err != nil {
		fmt.Printf("WARNING: webhook dispatch fetch endpoints: %v\n", err)
		return
	}
	for _, endpoint := range endpoints {
		go s.deliverWebhook(context.Background(), endpoint, event)
	}
}

// RetryFailedWebhooks — called by cron
func (s *paymentEngine) RetryFailedWebhooks(ctx context.Context) error {
	logs, err := s.paymentRepo.GetPendingWebhookRetries(ctx, time.Now())
	if err != nil {
		return fmt.Errorf("fetch pending retries: %w", err)
	}
	for _, log := range logs {
		endpoint, err := s.paymentRepo.GetWebhookEndpointByID(ctx, log.EndpointID)
		if err != nil {
			continue
		}
		var payload map[string]any
		json.Unmarshal([]byte(log.Payload), &payload)
		go s.deliverWebhook(context.Background(), endpoint, WebhookEvent{
			EventType: log.EventType,
			EventID:   log.EventID,
			Payload:   payload,
		})
	}
	return nil
}

func (s *paymentEngine) deliverWebhook(
	ctx context.Context,
	endpoint *models.WebhookEndpoint,
	event WebhookEvent,
) {
	body, _ := json.Marshal(map[string]any{
		"event_type": event.EventType,
		"event_id":   event.EventID,
		"timestamp":  time.Now().UTC(),
		"payload":    event.Payload,
	})

	signature := signWebhookPayload(body, endpoint.SigningSecret)

	log := &models.WebhookDeliveryLog{
		ID:          uuid.New(),
		TenantID:    endpoint.TenantID,
		EndpointID:  endpoint.ID,
		EventType:   event.EventType,
		EventID:     event.EventID,
		Payload:     string(body),
		Status:      "PENDING",
		AttemptNum:  1,
		AttemptedAt: time.Now(),
	}

	start := time.Now()
	statusCode, respBody, err := postWebhook(ctx, endpoint.URL, body, signature)
	latency := int(time.Since(start).Milliseconds())

	log.ResponseLatency = &latency
	log.ResponseCode = &statusCode
	log.ResponseBody = &respBody

	if err != nil || statusCode < 200 || statusCode >= 300 {
		log.Status = "FAILED"
		if err != nil {
			msg := err.Error()
			log.ErrorMessage = &msg
		}
		next := time.Now().Add(webhookRetryBackoff(log.AttemptNum))
		log.NextRetryAt = &next
	} else {
		log.Status = "SUCCESS"
	}

	_ = s.paymentRepo.CreateWebhookDeliveryLog(context.Background(), log)
}

func postWebhook(ctx context.Context, url string, body []byte, signature string) (int, string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return 0, "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Sharpnr-Signature", signature)

	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return resp.StatusCode, buf.String(), nil
}

func signWebhookPayload(body []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	return fmt.Sprintf("sha256=%x", mac.Sum(nil))
}

func webhookRetryBackoff(attempt int) time.Duration {
	backoffs := []time.Duration{
		1 * time.Minute,
		5 * time.Minute,
		15 * time.Minute,
		30 * time.Minute,
		60 * time.Minute,
	}
	if attempt-1 < len(backoffs) {
		return backoffs[attempt-1]
	}
	return 0
}

func generateSigningSecret() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
