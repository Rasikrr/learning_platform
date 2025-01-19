package http

import (
	"context"
	"fmt"
	"github.com/Rasikrr/learning_platform/configs"
	"github.com/Rasikrr/learning_platform/internal/ports/http/handlers/auth"
	authS "github.com/Rasikrr/learning_platform/internal/services/auth"
	"log"
	"net/http"
	"time"
)

const (
	name         = "http server"
	idleTimeout  = time.Minute * 3
	readTimeout  = time.Second * 60
	writeTimeout = time.Second * 60
)

type Server struct {
	name string
	port string
	host string
	srv  *http.Server
}

func NewServer(
	cfg *configs.Config,
	authService authS.Service,
) *Server {
	authController := auth.NewController(authService)

	router := http.NewServeMux()
	authController.Init(router)

	srv := &http.Server{
		Addr:         address(cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
	return &Server{
		name: name,
		port: cfg.Server.Port,
		host: cfg.Server.Host,
		srv:  srv,
	}
}

func address(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func (s *Server) Start() error {
	log.Println("starting http server")
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
