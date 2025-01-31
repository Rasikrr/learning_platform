package mail

import (
	"context"
	"github.com/Rasikrr/learning_platform/configs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient(t *testing.T) {
	ctx := context.Background()
	cfg := configs.Config{
		Mail: configs.MailConfig{
			Host:     "smtp.yandex.kz",
			Port:     587,
			User:     "rassulturtulov@yandex.kz",
			Password: "laqajhdfqownscxy",
			From:     "rassulturtulov@yandex.kz",
		},
	}
	c := NewClient(&cfg)
	err := c.Send(ctx, []string{"rs.t.95@mail.ru"}, "test", "test")
	require.NoError(t, err)
}
