package mail

import (
	"context"
	"crypto/tls"
	"github.com/go-gomail/gomail"
)

type Client interface {
	Send(ctx context.Context, from string, to []string, subject, body string) error
}

type client struct {
	srv *gomail.Dialer
}

func NewClient() Client {
	srv := gomail.NewDialer("smtp.yandex.ru", 587, "mail.denser.com", "rasik1234")
	srv.TLSConfig = &tls.Config{
		ServerName: "mail.denser.com",
		MinVersion: tls.VersionTLS12,
	}
	return &client{
		srv: srv,
	}
}

func (c *client) Send(_ context.Context, from string, to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return c.srv.DialAndSend(m)
}
