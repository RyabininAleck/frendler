package scheduler

import (
	"frendler/processor/scheduler/tasks"
)

type SchedulerImpl struct {
	tasks []tasks.Task
}

func (t *SchedulerImpl) StartDataTasks() {
	for _, task := range t.tasks {
		task.Run()
	}
}

func (t *SchedulerImpl) StopDataTasks() {
	for _, task := range t.tasks {
		task.Stop()
	}
}
