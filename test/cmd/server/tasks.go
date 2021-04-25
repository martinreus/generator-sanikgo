package server

import (
	"context"
	"github.com/apex/log"
	"sync"
	"test/pkg/tasks"
)

type configuration struct {
	taskList []tasks.Task
	tasksInfo *tasks.TaskInfoList
	m sync.Mutex
}

var Instance = configuration{
	taskList:  []tasks.Task{},
	tasksInfo: tasks.NewInfoList(),
	m:         sync.Mutex{},
}

func (c *configuration) AppendTask(task tasks.Task) {
	c.m.Lock()
	defer c.m.Unlock()

	c.taskList = append(c.taskList, task)
	c.tasksInfo.Append(task)
}

func (c *configuration) StartAllTasks(ctx context.Context) {
	// then start all of them
	for _, task := range c.taskList {
		task := task
		go func() {
			if err := task.Start(ctx); err != nil {
				log.WithError(err).Errorf("Error while starting task '%s'", task.Name())
			} else {
				log.Infof("Started task '%s'", task.Name())
			}
		}()
	}
}

func (c *configuration) StopAllTasks(ctx context.Context) {
	// stop all tasks
	for _, task := range c.taskList {
		task := task
		if err := task.Stop(ctx); err != nil {
			log.WithError(err).Errorf("Error while stopping task '%s'", task.Name())
		} else {
			log.Infof("Stopped task '%s'", task.Name())
		}
	}
}