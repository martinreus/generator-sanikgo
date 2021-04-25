package main

import (
	"context"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"os"
	"os/signal"
	"test/cmd/server"
)

func main() {

	ctx := context.Background()
	log.SetHandler(json.New(os.Stdout))

	server.Instance.StartAllTasks(ctx)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	<-c

	server.Instance.StopAllTasks(ctx)


}
