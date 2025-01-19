package mail

import (
	"context"
	"crypto/tls"
	"github.com/Rasikrr/learning_platform/configs"
	"github.com/go-gomail/gomail"
)

type Client interface {
	Send(ctx context.Context, to []string, subject, body string) error
}

type client struct {
	srv  *gomail.Dialer
	from string
}

func NewClient(cfg *configs.Config) Client {
	srv := gomail.NewDialer(cfg.Mail.Host, cfg.Mail.Port, cfg.Mail.User, cfg.Mail.Password)
	srv.TLSConfig = &tls.Config{
		ServerName: cfg.Mail.Host,
		MinVersion: tls.VersionTLS12,
	}
	return &client{
		srv:  srv,
		from: cfg.Mail.From,
	}
}

func (c *client) Send(_ context.Context, to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", c.from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return c.srv.DialAndSend(m)
}
