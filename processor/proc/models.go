package proc

import (
	"frendler/processor/adapter"
	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/tasks"
)

type Processor struct {
	config  config.Config
	DB      db.DB
	Adapter adapter.Adapter
	Tasks   tasks.Task
}
