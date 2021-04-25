package main

import (
	"context"
	"github.com/apex/log"
	"os"
	"os/signal"
	"test/cmd/config"
	"test/pkg/tasks"
)

func main() {

	var tasksInfo *[]tasks.Info
	var taskList []tasks.Task
	ctx := context.Background()

	// create a bunch of tasks
	taskList = append(taskList, config.ConfigureRest(tasksInfo))

	// aggregate all tasksInfo into the pointer above


	// then start all of them
	for _, task := range taskList {
		task := task
		go func() {
			if err := task.Start(ctx); err != nil {
				log.WithError(err).Errorf("Error while start task '%s'", task.Name())
			} else {
				log.Infof("Started task '%s'", task.Name())
			}
		}()
	}

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	<-c


	// stop all tasks
	for _, task := range taskList {
		task := task
		if err := task.Stop(ctx); err != nil {
			log.WithError(err).Errorf("Error while stopping task '%s'", task.Name())
		} else {
			log.Infof("Stopped task '%s'", task.Name())
		}
	}

}
