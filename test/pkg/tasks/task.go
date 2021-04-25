package tasks

import (
	"context"
	"sync"
)

var (
	Running State = "running"
	Stopped State = "stopped"
	Error   State = "error"
)

type State string

type Status struct {
	State       State
	Err         error
}

type Info interface {
	Status() []Status
	Name() string
}

type Task interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Info
}

type TaskInfoList struct {
	infos []Info
	m sync.Mutex
}

func NewInfoList() *TaskInfoList {
	return &TaskInfoList{
		infos: []Info{},
	}
}

func (til *TaskInfoList) Append(info Info)  {
	til.m.Lock()
	defer til.m.Unlock()

	til.infos = append(til.infos, info)
}

func (til *TaskInfoList) GetTaskInfos() []Info {
	return til.infos
}