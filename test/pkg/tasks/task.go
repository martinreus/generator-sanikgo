package tasks

import "context"

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
