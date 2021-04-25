package config

import (
	"github.com/go-chi/chi/v5/middleware"
	"test/internal/web/rest"
	"test/pkg/tasks"
)

func ConfigureRest(tasksInfo *[]tasks.Info) tasks.Task {

	server := rest.New(
		rest.WithMiddleware(middleware.Logger),
		rest.WithTaskInfoList(tasksInfo),
	)

	// TODO: add additional configuration here, such as dependencies the webserver needs - other services for instance

	return server

}
