package config

import (
	"<%=moduleName%>/<%=genOutputPath%>"
	"<%=moduleName%>/pkg/tasks"
	"github.com/go-chi/chi/v5/middleware"
)

func Configure<%=openApiGenPackageUpper%>(tasksInfo *[]tasks.Info) tasks.Task {

	server := <%=openApiGenPackage%>.New(
		<%=openApiGenPackage%>.WithMiddleware(middleware.Logger),
		<%=openApiGenPackage%>.WithTaskInfoList(tasksInfo),
	)

	// TODO: additional configuration here, such as dependencies the webserver needs - other services for instance

	return server

}