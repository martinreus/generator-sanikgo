package <%=openApiGenPackage%>

import "<%=moduleName%>/pkg/env"

type Config struct {
	ServerPort int    `json:"serverPort""`
	BaseUrl    string `json:"baseUrl"`
}

func ConfigFromEnv() Config {
	return Config{
		ServerPort: env.GetIntOrDefault("<%=openApiGenPackageUpper%>_SERVER_PORT", <%=restApiPort%>),
		BaseUrl:    env.GetStringOrDefault("<%=openApiGenPackageUpper%>_BASE_URL", ""),
	}
}
