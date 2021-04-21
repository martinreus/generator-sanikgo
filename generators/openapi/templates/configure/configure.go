package config

import (
	"<%=moduleName%>/<%=genOutputPath%>"
	"<%=moduleName%>/pkg/tasks"
	"github.com/go-chi/chi/v5/middleware"
)

func configure<%=openApiGenPackage%>() tasks.Task {

	server := <%=openApiGenPackage%>.New(<%=openApiGenPackage%>.WithMiddleware(middleware.Logger))

	// TODO: additional configuration here, such as dependencies the webserver needs - other services for instance

	return server

}