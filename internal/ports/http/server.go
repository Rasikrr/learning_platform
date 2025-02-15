package http

// nolint: revive
import (
	"context"
	"fmt"
	"github.com/Rasikrr/learning_platform/configs"
	_ "github.com/Rasikrr/learning_platform/docs"
	adminAuth "github.com/Rasikrr/learning_platform/internal/ports/http/handlers/admin/auth"
	"github.com/Rasikrr/learning_platform/internal/ports/http/handlers/auth"
	"github.com/Rasikrr/learning_platform/internal/ports/http/handlers/courses/commands"
	"github.com/Rasikrr/learning_platform/internal/ports/http/handlers/courses/queries"
	"github.com/Rasikrr/learning_platform/internal/ports/http/handlers/enrollments"
	"github.com/Rasikrr/learning_platform/internal/ports/http/handlers/faq"
	"github.com/Rasikrr/learning_platform/internal/ports/http/handlers/users"
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	authS "github.com/Rasikrr/learning_platform/internal/services/auth"
	coursesS "github.com/Rasikrr/learning_platform/internal/services/courses"
	enrollmentsS "github.com/Rasikrr/learning_platform/internal/services/enrollments"
	faqS "github.com/Rasikrr/learning_platform/internal/services/faq"
	filesS "github.com/Rasikrr/learning_platform/internal/services/files"
	submissionS "github.com/Rasikrr/learning_platform/internal/services/submissions"
	usersS "github.com/Rasikrr/learning_platform/internal/services/users"
	httpSwagger "github.com/swaggo/http-swagger"
	"os"

	"log"
	"net/http"
	"time"
)

const (
	name         = "http server"
	idleTimeout  = time.Minute * 3
	readTimeout  = time.Second * 60
	writeTimeout = time.Second * 60
	fileDir      = "./files"
	avatarDir    = fileDir + "/avatars"
	courseDir    = fileDir + "/courses"
	faqDir       = fileDir + "/faq"
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

// @host localhost:8081
// @BasePath /

// Server @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type Server struct {
	name string
	port string
	host string
	srv  *http.Server
}

func NewServer(
	cfg *configs.Config,
	authService authS.Service,
	courseService coursesS.Service,
	enrollmentsService enrollmentsS.Service,
	submissionsService submissionS.Service,
	faqService faqS.Service,
	usersService usersS.Service,
	filesService filesS.Service,
) *Server {
	router := http.NewServeMux()

	// Middlewares
	authMiddleware := middlewares.NewAuthMiddleware(authService)
	enrollmentMiddleware := middlewares.NewEnrollMiddleware(enrollmentsService)

	// Controllers
	faqController := faq.NewController(authMiddleware, faqService)
	authController := auth.NewController(authService)
	coursesQueriesController := queries.NewController(courseService, authMiddleware, enrollmentMiddleware)
	enrollmentsController := enrollments.NewController(enrollmentsService, authMiddleware)
	courseCommandsController := commands.NewController(courseService, authMiddleware, enrollmentMiddleware, submissionsService)
	usersController := users.NewController(usersService, filesService, authMiddleware)
	adminAuthController := adminAuth.NewController(authService)

	// Init controllers
	coursesQueriesController.Init(router)
	authController.Init(router)
	faqController.Init(router)
	enrollmentsController.Init(router)
	courseCommandsController.Init(router)
	usersController.Init(router)
	adminAuthController.Init(router)

	// CORS
	r := middlewares.CORSMiddleware(router)
	r = middlewares.PanicHandler(r)

	srv := &http.Server{
		Addr:         address(cfg.Server.Host, cfg.Server.Port),
		Handler:      r,
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
	if err := s.initStatic(router); err != nil {
		log.Fatal(err)
	}
	return s
}

func address(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func (s *Server) initSwagger(router *http.ServeMux) {
	addr := hostPort(s.host, s.port)
	url := fmt.Sprintf("http://%s/swagger/doc.json", addr)
	router.Handle("/swagger/",
		httpSwagger.Handler(httpSwagger.URL(url)),
	)
}

// nolint: unparam
func (s *Server) initStatic(router *http.ServeMux) error {
	dirs := []string{fileDir, avatarDir, courseDir, faqDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	router.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	router.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./files/avatars"))))
	router.Handle("/courses/", http.StripPrefix("/courses/", http.FileServer(http.Dir("./files/courses"))))
	router.Handle("/faq/", http.StripPrefix("/faq/", http.FileServer(http.Dir("./files/faq"))))
	return nil
}

func (s *Server) Start() error {
	log.Println("starting http server")
	log.Printf("swagger url: http://%s/swagger/index.html\n", hostPort(s.host, s.port))
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func hostPort(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
