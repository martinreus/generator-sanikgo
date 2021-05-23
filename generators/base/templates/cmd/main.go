package main

import (
	"context"
	"os"
	"os/signal"
	"<%=moduleName%>/cmd/app"
	"<%=moduleName%>/pkg/logging"
)

func main() {

	ctx := context.Background()

	logging.Configure()
	app.Instance.StartAllTasks(ctx)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	<-c

	app.Instance.StopAllTasks(ctx)


}
