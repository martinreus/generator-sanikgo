package <%=openApiGenPackage%>

import "test/pkg/env"

type Config struct {
	ServerPort int
}

func ConfigFromEnv() Config {
	return Config{
		ServerPort: env.GetIntOrDefault("<%=openApiGenPackageUpper%>_SERVER_PORT", <%=restApiPort%>),
	}
}