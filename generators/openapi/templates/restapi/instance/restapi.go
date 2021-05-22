package app

import (
	"github.com/go-chi/chi/v5/middleware"
	"<%=moduleName%>/<%=genOutputPath%>"
)

func init() {
	webServerInstance := <%=openApiGenPackage%>.New(
		<%=openApiGenPackage%>.WithMiddleware(middleware.Logger),
		<%=openApiGenPackage%>.WithTaskInfoList(Instance.tasksInfo),
		<%=openApiGenPackage%>.WithConfig(restapi.ConfigFromEnv()),
	)

	// TODO: add additional dependencies, like database connections, etc...

	Instance.AppendTask(webServerInstance)
}

