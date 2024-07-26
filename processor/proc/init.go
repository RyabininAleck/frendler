package proc

import (
	"frendler/processor/adapter"
	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/proc/handler"
	"frendler/processor/scheduler"
)

func Init(config *config.Config, DB db.DB, Handler handler.Handler, Adapter adapter.Adapter, Tasks scheduler.Scheduler) Processor {
	return Processor{
		config:    config,
		Handler:   Handler,
		DB:        DB,
		Adapter:   Adapter,
		Scheduler: Tasks,
	}
}
