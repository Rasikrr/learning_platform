package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/Rasikrr/learning_platform/configs"
	"github.com/Rasikrr/learning_platform/internal/databases"
	http "github.com/Rasikrr/learning_platform/internal/ports/http"
	usersR "github.com/Rasikrr/learning_platform/internal/repositories/users"
	usersS "github.com/Rasikrr/learning_platform/internal/services/users"
	"github.com/Rasikrr/learning_platform/internal/workers"
	dbInfo "github.com/Rasikrr/learning_platform/internal/workers/db_info"
	"github.com/hashicorp/go-multierror"
	"log"
	HTTP "net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Starter interface {
	Start() error
}

type App struct {
	name     string
	config   configs.Config
	postgres *databases.Postgres

	workers         []workers.Worker
	usersRepository usersR.Repository
	userService     usersS.Service
	httpServer      *http.Server
}

func InitApp(ctx context.Context, name string) *App {
	cfg, err := configs.Parse()
	if err != nil {
		panic(err)
	}
	app := &App{
		name:   name,
		config: cfg,
	}
	for _, init := range []func(ctx context.Context) error{
		app.InitPostgres,
		app.InitRepositories,
		app.InitServices,
		app.InitHTTPServer,
		app.InitWorkers,
	} {
		if err := init(ctx); err != nil {
			log.Fatal(ctx, "init app", err)
		}
	}
	return app
}

func (a *App) InitRepositories(_ context.Context) error {
	a.usersRepository = usersR.NewRepository()
	return nil
}

func (a *App) InitServices(_ context.Context) error {
	a.userService = usersS.NewService(a.usersRepository)
	return nil
}

func (a *App) InitHTTPServer(_ context.Context) error {
	a.httpServer = http.NewServer(&a.config, a.userService)
	return nil
}

func (a *App) InitWorkers(_ context.Context) error {
	a.workers = []workers.Worker{
		dbInfo.NewWorker(a.usersRepository),
	}
	return nil
}

func (a *App) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	stop := make(chan struct{})

	go a.handleShutdown(cancel, stop)

	fns := make([]any, 0, len(a.workers)+1)
	fns = append(fns, a.httpServer.Start)
	for _, w := range a.workers {
		fns = append(fns, w.Run)
	}

	if err := runParallel(ctx, fns...); err != nil {
		return err
	}
	<-stop
	return nil
}

func (a *App) handleShutdown(cancel context.CancelFunc, s chan struct{}) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	log.Println("Received shutdown signal")

	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		log.Println("Error while shutting down HTTP server:", err)
	}
	cancel()
	close(s)
}

func (a *App) InitPostgres(ctx context.Context) error {
	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		"postgres", "postgres", "localhost", "5432", "postgres", "sslmode=disable")
	a.postgres, err = databases.NewPostgres(ctx, &a.config, dsn)
	if err != nil {
		return err
	}
	log.Println("Postgres connected")
	return nil
}

func runParallel(ctx context.Context, fns ...any) error {
	var (
		wg     sync.WaitGroup
		errRes *multierror.Error
	)
	for _, f := range fns {
		wg.Add(1)
		go func(fn any) {
			defer wg.Done()
			switch fun := fn.(type) {
			case func() error:
				if err := fun(); err != nil {
					if errors.Is(err, HTTP.ErrServerClosed) {
						log.Println("HTTP server closed gracefully")
						return
					}
					errRes = multierror.Append(errRes, err)
				}
			case func(ctx context.Context) error:
				if err := fun(ctx); err != nil {
					errRes = multierror.Append(errRes, err)
				}
			default:
				log.Printf("unknown function type: %T", fn)
			}
		}(f)
	}
	wg.Wait()
	return errRes.ErrorOrNil()
}
