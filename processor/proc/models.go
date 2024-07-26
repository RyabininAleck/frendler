package proc

import (
	"frendler/processor/adapter"
	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/proc/handler"
	"frendler/processor/scheduler"
)

type Processor struct {
	config    *config.Config
	Handler   handler.Handler
	DB        db.DB
	Adapter   adapter.Adapter
	Scheduler scheduler.Scheduler
}
