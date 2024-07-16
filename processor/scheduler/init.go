package scheduler

import (
	"context"
	"time"

	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/scheduler/tasks"
)

func Init(cfg config.TaskConf, db db.DB) *SchedulerImpl {
	var ts []tasks.Task

	ctx, cancel := context.WithCancel(context.Background())
	contactUpdateTask := tasks.CreateContactUpdateTask(ctx, cancel, time.Duration(cfg.Interval)*time.Second, db)
	ts = append(ts, contactUpdateTask)

	//todo тут добавляются таски, по аналогии contactUpdateTask
	return &SchedulerImpl{
		tasks: ts,
	}
}
