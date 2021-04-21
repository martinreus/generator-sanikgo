package tasks

import "context"

type Task interface {
	// Start implementations should start a non blocking go-routine as soon as possible
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	TaskName() string
}
