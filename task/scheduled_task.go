package task

import (
	"github.com/procyon-projects/chrono"
)

type ScheduledTask interface {
	Name() string
	Run() error
	Cancel()
}

var TaskScheduler = chrono.NewDefaultTaskScheduler()
