package main

import (
	"context"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"os"
	"os/signal"
	"test/cmd/app"
)

func main() {

	ctx := context.Background()
	log.SetHandler(json.New(os.Stdout))

	app.Instance.StartAllTasks(ctx)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	<-c

	app.Instance.StopAllTasks(ctx)


}
