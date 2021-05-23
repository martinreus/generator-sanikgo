package app

import (
	"<%=moduleName%>/<%=genOutputPath%>"
	"<%=moduleName%>/pkg/middlewares"
)

func init() {
	webServerInstance := <%=openApiGenPackage%>.New(
		<%=openApiGenPackage%>.WithMiddleware(middlewares.Recover),
		<%=openApiGenPackage%>.WithMiddleware(middlewares.Logger()),
		<%=openApiGenPackage%>.WithTaskInfoList(Instance.tasksInfo),
		<%=openApiGenPackage%>.WithConfig(restapi.ConfigFromEnv()),
	)

	// TODO: add additional dependencies, like database connections, etc...

	Instance.AppendTask(webServerInstance)
}

