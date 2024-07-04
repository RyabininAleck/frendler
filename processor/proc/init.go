package proc

import (
	"frendler/processor/adapter"
	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/proc/handler"
	"frendler/processor/tasks"
)

func Init(config config.Config, DB db.DB, Handler handler.Handler, Adapter adapter.Adapter, Tasks tasks.Task) Processor {
	return Processor{
		config:  config,
		Handler: Handler,
		DB:      DB,
		Adapter: Adapter,
		Tasks:   Tasks,
	}
}
