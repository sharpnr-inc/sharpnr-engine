// GatewayResolver bridges PaymentEngineService and the ProviderRegistry.
// Responsible for:
//   1. Fetching the correct PaymentGatewayConfig from DB for tenant + provider
//   2. Decrypting the stored credentials
//   3. Injecting the WebhookURL into the config map
//   4. Returning the correct GatewayProvider from the registry
//
// RESOLUTION PRIORITY (legal entity aware):
//   1. TenantID + LegalEntityID + ProviderCode  (legal entity specific config)
//   2. TenantID + nil LegalEntityID + ProviderCode  (tenant-wide fallback)
//
// WHY LEGAL ENTITY MATTERS:
//   A tenant can have multiple legal entities (Maharashtra + Delhi branches).
//   Each legal entity may have its own PhonePe MerchantID and credentials.
//   The resolver ensures the correct credentials are used per transaction.
//
// WEBHOOK URL INJECTION:
//   WebhookURL is stored once in PaymentGatewayConfig.WebhookURL during gateway setup.
//   The resolver injects it into the Config map as "webhook_url".
//   Each gateway provider reads it from Config — no per-transaction URL passing needed.

package payment_gateway

import (
	"context"
	"encoding/json"
	"fmt"

	models "sharpnr-engine/payment_engine/models"
	repo "sharpnr-engine/payment_engine/repo"

	"github.com/google/uuid"
)

// ResolvedGateway is the output of the resolver.
// Contains everything needed to make a gateway call.
type ResolvedGateway struct {
	Provider GatewayProvider
	Config   map[string]string // decrypted config + webhook URL
	ConfigID uuid.UUID
}

type GatewayResolver struct {
	repo     repo.PaymentEngineRepoInterface
	registry *ProviderRegistry
}

func NewGatewayResolver(
	repo repo.PaymentEngineRepoInterface,
	registry *ProviderRegistry,
) *GatewayResolver {
	return &GatewayResolver{
		repo:     repo,
		registry: registry,
	}
}

// ResolveByProvider resolves gateway by tenant + legal entity + provider code.
// Called by PaymentService.Pay() when initiating a new payment.
func (r *GatewayResolver) ResolveByProvider(
	ctx context.Context,
	tenantID uuid.UUID,
	legalEntityID uuid.UUID,
	providerCode string,
) (*ResolvedGateway, error) {

	// 1. Try legal entity specific config first
	config, err := r.repo.GetActiveGatewayConfig(ctx, tenantID,  providerCode)
	if err != nil {
		// 2. Fall back to tenant-wide config (LegalEntityID = nil)
		config, err = r.repo.GetActiveGatewayConfig(ctx, tenantID, providerCode)
		if err != nil {
			return nil, fmt.Errorf(
				"no active gateway config found for provider '%s' (tenant: %s, legalEntity: %s)",
				providerCode, tenantID, legalEntityID,
			)
		}
	}

	return r.build(config)
}

// ResolveByConfigID resolves gateway by a known config ID.
// Called by:
//   - PaymentService.ConfirmPayment() — already has GatewayConfigID from Payment record
//   - RefundService.InitiateRefund() — uses config from original payment
func (r *GatewayResolver) ResolveByConfigID(
	ctx context.Context,
	configID uuid.UUID,
) (*ResolvedGateway, error) {

	config, err := r.repo.GetGatewayConfigByID(ctx, configID)
	if err != nil {
		return nil, fmt.Errorf("gateway config not found: %w", err)
	}

	return r.build(config)
}

// build decrypts config, injects webhook URL, resolves provider.
func (r *GatewayResolver) build(config *models.PaymentGatewayConfig) (*ResolvedGateway, error) {

	// Decrypt stored credentials
	decrypted, err := decryptConfig(config.ConfigEncrypted)
	if err != nil {
		return nil, fmt.Errorf("decrypt gateway config: %w", err)
	}

	// Inject webhook URL — gateway provider reads from config map
	// This URL is registered on PhonePe/Razorpay dashboard by tenant during setup
	decrypted["webhook_url"] = config.WebhookURL
	decrypted["is_test_mode"] = boolToString(config.IsTestMode)

	// Resolve provider from registry
	provider, err := r.registry.Resolve(config.ProviderCode)
	if err != nil {
		return nil, err
	}

	return &ResolvedGateway{
		Provider: provider,
		Config:   decrypted,
		ConfigID: config.ID,
	}, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// ENCRYPTION HELPERS
// Replace with your actual AES/KMS implementation in production.
// The interface stays the same — only the internals change.
// ─────────────────────────────────────────────────────────────────────────────

// DecryptConfig decrypts the stored config JSON.
// TODO: Replace json.Unmarshal with AES-256-GCM or KMS decryption.
func decryptConfig(encrypted string) (map[string]string, error) {
	var result map[string]string
	if err := json.Unmarshal([]byte(encrypted), &result); err != nil {
		return nil, fmt.Errorf("decryptConfig: %w", err)
	}
	return result, nil
}

// EncryptConfig encrypts config before saving to DB.
// TODO: Replace json.Marshal with AES-256-GCM or KMS encryption.
func EncryptConfig(config map[string]string) (string, error) {
	raw, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
