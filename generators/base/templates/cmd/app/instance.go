package app

import (
	"context"
	"github.com/apex/log"
	"sync"
	"<%=moduleName%>/pkg/tasks"
)

type serverInstance struct {
	taskList []tasks.Task
	tasksInfo *tasks.TaskInfoList
	m sync.Mutex
}

var Instance = serverInstance{
	taskList:  []tasks.Task{},
	tasksInfo: tasks.NewInfoList(),
	m:         sync.Mutex{},
}

func (c *serverInstance) AppendTask(task tasks.Task) {
	c.m.Lock()
	defer c.m.Unlock()

	c.taskList = append(c.taskList, task)
	c.tasksInfo.Append(task)
}

func (c *serverInstance) StartAllTasks(ctx context.Context) {
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

func (c *serverInstance) StopAllTasks(ctx context.Context) {
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
