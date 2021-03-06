package <%=openApiGenPackage%>

import (
	"context"
	"errors"
	"fmt"
	"github.com/apex/log"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sync"
	"<%=moduleName%>/pkg/tasks"
	"time"
)

type ServerOption func(instance *serverInstance)

type serverInstance struct {
	m           sync.Mutex
	middlewares []func(http.Handler) http.Handler
	server      *http.Server
	tasksInfo   *tasks.TaskInfoList
	status      tasks.Status
	config      Config
}

func WithConfig(config Config) ServerOption {
	return func(instance *serverInstance) {
		instance.config = config
	}
}

func WithTaskInfoList(tasksInfo *tasks.TaskInfoList) ServerOption {
	return func(instance *serverInstance) {
		instance.tasksInfo = tasksInfo
	}
}

func WithMiddleware(middleware func(http.Handler) http.Handler) ServerOption {
	return func(instance *serverInstance) {
		instance.middlewares = append(instance.middlewares, middleware)
	}
}

func New(serverOpts ...ServerOption) *serverInstance {
	s := &serverInstance{
		status: tasks.Status{
			State: tasks.Stopped,
			Err:   errors.New("webserver not running"),
		},
		config: Config{
			ServerPort: 8080,
		},
	}

	for _, opt := range serverOpts {
		opt(s)
	}
	return s
}

func (s *serverInstance) Start(ctx context.Context) error {
	log.WithFields(log.Fields{
		"serverPort": s.config.ServerPort,
		"baseUrl": s.config.BaseUrl,
	}).Infof("Server starting..")

	router := chi.NewRouter()

	var middlewares []MiddlewareFunc

	for _, middleware := range s.middlewares {
		middleware := middleware
		middlewares = append(middlewares, func(han http.HandlerFunc) http.HandlerFunc {
			return middleware(han).ServeHTTP
		})
	}

	router.Mount("/", HandlerWithOptions(
		s, ChiServerOptions{
			BaseURL:     s.config.BaseUrl,
			Middlewares: middlewares,
		}))

	s.server = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", s.config.ServerPort),
		Handler: router,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	s.server.RegisterOnShutdown(func() {
		s.m.Lock()
		defer s.m.Unlock()
		s.status = tasks.Status{
			State: tasks.Stopped,
			Err:   nil,
		}
	})

	go func() {
		s.m.Lock()
		s.status = tasks.Status{
			State: tasks.Running,
			Err:   nil,
		}
		s.m.Unlock()
		_ = s.server.ListenAndServe()
	}()
	return nil
}

func (s *serverInstance) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *serverInstance) Name() string {
	return "<%=openApiGenPackage%>"
}

func (s *serverInstance) Status() []tasks.Status {
	s.m.Lock()
	defer s.m.Unlock()
	return []tasks.Status{s.status}
}
