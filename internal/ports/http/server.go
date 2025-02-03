package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/Rasikrr/learning_platform/configs"
	_ "github.com/Rasikrr/learning_platform/docs"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/ports/http/handlers/auth"
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	authS "github.com/Rasikrr/learning_platform/internal/services/auth"
	httpSwagger "github.com/swaggo/http-swagger"
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

// @title Learning Platform API
// @version 1.0
// @description This is docs for Learning Platform API
// @termsOfService http://swagger.io/terms/

// @contact.name Pick me team
// @contact.url https://github.com/Rasikrr/learning_platform
// @contact.email leaning_platform@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api-golang-production-a90c.up.railway.app
// @BasePath /api/v1
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
	// authMiddleware := middlewares.NewAuthMiddleware(authService)

	authController := auth.NewController(authService)
	router := http.NewServeMux()
	authController.Init(router)

	routerWithCORS := middlewares.CORSMiddleware(router)

	srv := &http.Server{
		Addr:         address(cfg.Server.Host, cfg.Server.Port),
		Handler:      routerWithCORS,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	s := &Server{
		name: name,
		port: cfg.Server.Port,
		host: cfg.Server.Host,
		srv:  srv,
	}
	s.initSwagger(router)
	return s
}

func address(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func (s *Server) initSwagger(router *http.ServeMux) {
	hostPort := fmt.Sprintf("%s:%s", s.host, s.port)
	url := fmt.Sprintf("http://%s/swagger/doc.json", hostPort)
	router.Handle("/swagger/",
		httpSwagger.Handler(httpSwagger.URL(url)),
	)
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

func GetSession(ctx context.Context) (*entity.Session, error) {
	token := ctx.Value(middlewares.SessionKey)
	if token == nil {
		return nil, errors.New("session is empty")
	}
	s, ok := token.(*entity.Session)
	if !ok {
		return nil, errors.New("session is not Session")
	}
	return s, nil
}
