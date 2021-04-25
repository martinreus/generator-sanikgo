package server

import (
	"github.com/go-chi/chi/v5/middleware"
	"test/internal/web/rest"
)

func init() {
	restApiInstance := rest.New(
		rest.WithMiddleware(middleware.Logger),
		rest.WithTaskInfoList(Instance.tasksInfo),
	)

	// TODO: add additional dependencies, like database connections, etc...

	Instance.AppendTask(restApiInstance)
}
