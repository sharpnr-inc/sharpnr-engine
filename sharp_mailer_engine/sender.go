package sharp_mailer_engine

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"math/big"
	"net/smtp"
	"time"
)

type smtpClientWrapper struct {
	smtp *smtp.Client
}

func (c *smtpClientWrapper) IsClosed() bool {
	// A simple Noop check to see if the server hung up
	return c.smtp.Noop() != nil
}

func (e *SharpMailerEngine) dialNewClient() (*smtpClientWrapper, error) {
	s := e.cfg.SMTP
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	tlsConfig := &tls.Config{ServerName: s.Host}
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		return nil, err
	}

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	if err = client.Auth(auth); err != nil {
		return nil, err
	}

	return &smtpClientWrapper{smtp: client}, nil
}

func (e *SharpMailerEngine) sendPooled(c *smtp.Client, to, subject, htmlBody, from string) error {
	if err := c.Mail(from); err != nil {
		return err
	}
	if err := c.Rcpt(to); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	msgId, _ := generateMessageID(e.cfg.SMTP.Host)
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMessage-ID: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\nDate: %s\r\n\r\n%s",
		from, to, subject, msgId, time.Now().Format(time.RFC1123Z), htmlBody)

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	return w.Close()
}

// generateMessageID creates a unique ID to help with email deliverability.
func generateMessageID(host string) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("<%d.%d@%s>", time.Now().UnixNano(), n, host), nil
}
