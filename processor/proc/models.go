package proc

import (
	"frendler/processor/adapter"
	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/proc/handler"
	"frendler/processor/tasks"
)

type Processor struct {
	config  config.Config
	Handler handler.Handler
	DB      db.DB
	Adapter adapter.Adapter
	Tasks   tasks.Task
}
