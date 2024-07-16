package scheduler

import (
	"context"
	"time"

	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/scheduler/tasks"
)

func Init(ctx context.Context, cfg config.TaskConf, db db.DB) *SchedulerImpl {
	var ts []tasks.Task

	contactUpdateTask := tasks.CreateContactUpdateTask(ctx, time.Duration(cfg.Interval)*time.Second, db)
	ts = append(ts, contactUpdateTask)

	return &SchedulerImpl{
		tasks: ts,
	}
}
