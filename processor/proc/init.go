package proc

import (
	"frendler/processor/adapter"
	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/tasks"
)

func Init(config config.Config, DB db.DB, Adapter adapter.Adapter, Tasks tasks.Task) Processor {
	return Processor{
		config:  config,
		DB:      DB,
		Adapter: Adapter,
		Tasks:   Tasks,
	}
}
