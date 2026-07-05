// sender.go

package sharp_mailer_engine

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"math/big"
	"mime"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type smtpClientWrapper struct {
	smtp *smtp.Client
	conn net.Conn
}

func (c *smtpClientWrapper) IsClosed() bool {
	// A simple Noop check to see if the server hung up.
	// Deadline keeps a dead-but-open connection from blocking the worker.
	c.conn.SetDeadline(time.Now().Add(10 * time.Second))
	defer c.conn.SetDeadline(time.Time{})
	return c.smtp.Noop() != nil
}

func (e *SharpMailerEngine) dialNewClient() (*smtpClientWrapper, error) {
	s := e.cfg.SMTP
	addr := net.JoinHostPort(s.Host, fmt.Sprintf("%d", s.Port))
	tlsConfig := &tls.Config{ServerName: s.Host}
	dialer := &net.Dialer{Timeout: s.ConnectTimeout}

	var conn net.Conn
	var client *smtp.Client
	var err error

	if s.Encryption == "ssl" || s.Port == 465 {
		// Implicit TLS (SMTPS, port 465)
		conn, err = tls.DialWithDialer(dialer, "tcp", addr, tlsConfig)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(s.ConnectTimeout))
		client, err = smtp.NewClient(conn, s.Host)
		if err != nil {
			conn.Close()
			return nil, err
		}
	} else {
		// Plain dial then STARTTLS (port 587)
		conn, err = dialer.Dial("tcp", addr)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(s.ConnectTimeout))
		client, err = smtp.NewClient(conn, s.Host)
		if err != nil {
			conn.Close()
			return nil, err
		}
		if err = client.StartTLS(tlsConfig); err != nil {
			client.Close()
			return nil, err
		}
	}

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	if err = client.Auth(auth); err != nil {
		client.Close()
		return nil, err
	}

	conn.SetDeadline(time.Time{})
	return &smtpClientWrapper{smtp: client, conn: conn}, nil
}

func (e *SharpMailerEngine) sendPooled(c *smtpClientWrapper, to, subject, htmlBody, from string) error {
	c.conn.SetDeadline(time.Now().Add(e.cfg.SMTP.SendTimeout))
	defer c.conn.SetDeadline(time.Time{})

	to = sanitizeHeader(to)
	from = sanitizeHeader(from)

	if err := c.smtp.Mail(from); err != nil {
		return err
	}
	if err := c.smtp.Rcpt(to); err != nil {
		return err
	}

	w, err := c.smtp.Data()
	if err != nil {
		return err
	}

	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMessage-ID: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\nDate: %s\r\n\r\n%s",
		from, to, encodeSubject(subject), generateMessageID(from), time.Now().Format(time.RFC1123Z), htmlBody)

	if _, err = w.Write([]byte(message)); err != nil {
		return err
	}

	return w.Close()
}

// sanitizeHeader strips CR/LF so user-supplied values (recipient
// address, subject) can't inject extra SMTP headers.
func sanitizeHeader(v string) string {
	return strings.NewReplacer("\r", "", "\n", "").Replace(v)
}

// encodeSubject RFC 2047-encodes subjects containing non-ASCII;
// plain ASCII subjects pass through unchanged.
func encodeSubject(subject string) string {
	return mime.QEncoding.Encode("UTF-8", sanitizeHeader(subject))
}

// generateMessageID creates a unique ID scoped to the sending domain
// to help with email deliverability.
func generateMessageID(from string) string {
	domain := strings.TrimSuffix(from[strings.LastIndex(from, "@")+1:], ">")
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return fmt.Sprintf("<%d@%s>", time.Now().UnixNano(), domain)
	}
	return fmt.Sprintf("<%d.%d@%s>", time.Now().UnixNano(), n, domain)
}
