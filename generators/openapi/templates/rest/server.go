package  <%=openApiGenPackage%>

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type ServerOption func(instance *serverInstance)

type serverInstance struct {
	m           sync.Mutex
	middlewares []func(http.Handler) http.Handler
	baseUrl     string
	logger      *log.Logger
	server      *http.Server
}

func WithMiddleware(middleware func(http.Handler) http.Handler) ServerOption {
	return func(instance *serverInstance) {
		instance.middlewares = append(instance.middlewares, middleware)
	}
}

func WithBaseUrl(baseUrl string) ServerOption {
	return func(instance *serverInstance) {
		instance.baseUrl = baseUrl
	}
}

func WithLogger(logger *log.Logger) ServerOption {
	return func(instance *serverInstance) {
		instance.logger = logger
	}
}

func New(serverOpts ...ServerOption) *serverInstance {
	s := &serverInstance{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}

	for _, opt := range serverOpts {
		opt(s)
	}
	return s
}

func (s *serverInstance) Start(ctx context.Context) error {
	router := chi.NewRouter()

	var middlewares []MiddlewareFunc

	for _, middleware := range s.middlewares {
		middlewares = append(middlewares, func(han http.HandlerFunc) http.HandlerFunc {
			return middleware(han).ServeHTTP
		})
	}

	router.Mount("/", HandlerWithOptions(
		s, ChiServerOptions{
			BaseURL:     s.baseUrl,
			Middlewares: middlewares,
		}))

	s.server = &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		s.logger.Print("Starting webserver..")
		_ = s.server.ListenAndServe()
	}()
	return nil
}

func (s *serverInstance) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *serverInstance) TaskName() string {
	return "<%=openApiGenPackage%>"
}
