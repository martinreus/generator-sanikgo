package <%=openApiGenPackage%>

import (
	"encoding/json"
	"net/http"
)

func (s *serverInstance) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	// TODO: if required, create a more complex health status depending on other systems you depend on, such as DB being reachable ;)
	json.NewEncoder(w).Encode(HealthStatus{
		Status: HealthStatusStatusOK,
	})
}