package <%=openApiGenPackage%>

import (
	"encoding/json"
	"fmt"
	"net/http"
	"<%=moduleName%>/pkg/ptrs"
	"<%=moduleName%>/pkg/tasks"
)

func (s *serverInstance) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	// iterate over all tasks running on this server
	errorList := s.createErrorListFromTasks()
	healthStatus := HealthStatusStatusOK
	if len(errorList) > 0 {
		healthStatus = HealthStatusStatusUNAVAILABLE
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	json.NewEncoder(w).Encode(HealthStatus{
		Status: healthStatus,
		Errors: &errorList,
	})
}

func (s *serverInstance) createErrorListFromTasks() []Error {
	var errorList []Error
	for _, taskInfo := range s.tasksInfo.GetTaskInfos() {
		for _, status := range taskInfo.Status() {
			if status.State == tasks.Error {
				errorList = append(errorList, Error{
					Code:    TaskError,
					Message: ptrs.Str(fmt.Sprintf("Task '%s' contains errors: %v", taskInfo.Name(), status.Err.Error())),
				})
			}
		}
	}
	return errorList
}
