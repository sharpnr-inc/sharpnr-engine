package sharp_mailer_engine

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"math/big"
	"net/smtp"
	"time"
)

func (e *SharpMailerEngine) send(to, subject, htmlBody, fromAlias string) error {
	s := e.cfg.SMTP
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	// 1. Establish a Secure TLS Connection FIRST (Fixes EOF)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         s.Host,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", addr, err)
	}
	defer conn.Close()

	// 2. Create the SMTP Client from the secure connection
	client, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		return fmt.Errorf("failed to create smtp client: %w", err)
	}
	defer client.Quit()

	// 3. Authenticate
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("auth failed: %w", err)
	}

	// 4. Set the "Envelope" (This makes Aliases work)
	if err = client.Mail(fromAlias); err != nil {
		return fmt.Errorf("failed to set MAIL FROM: %w", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set RCPT TO: %w", err)
	}

	// 5. Send the Data
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to open data writer: %w", err)
	}

	msgId, _ := generateMessageID(s.Host)
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMessage-ID: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		fromAlias, to, subject, msgId, htmlBody)

	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	return nil
}

// generateMessageID creates a unique ID to help with email deliverability.
func generateMessageID(host string) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("<%d.%d@%s>", time.Now().UnixNano(), n, host), nil
}
