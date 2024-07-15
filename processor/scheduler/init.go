package scheduler

import (
	"time"

	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/scheduler/tasks"
)

func Init(cfg config.TaskConf, db db.DB) *SchedulerImpl {
	var ts []tasks.Task

	contactUpdateTask := tasks.CreateContactUpdateTask(db, time.Duration(cfg.Interval)*time.Second)
	ts = append(ts, contactUpdateTask)

	return &SchedulerImpl{
		tasks: ts,
	}
}
