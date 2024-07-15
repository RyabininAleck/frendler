package main

import (
	"frendler/processor/adapter"
	"frendler/processor/config"
	"frendler/processor/db"
	"frendler/processor/proc"
	"frendler/processor/proc/handler"
	"frendler/processor/scheduler"
	"frendler/storage"
)

func main() {

	cfg := config.Get()

	procDB := db.Init(cfg.DB)
	storage.Migrations(procDB)

	procHandler := handler.Init(procDB)

	// todo подключиться к адаптеру, настроить
	procAdapter := adapter.Init(cfg.Adapter)
	// todo подключиться дататаски
	procScheduler := scheduler.Init(cfg.Task, procDB)

	processor := proc.Init(cfg, procDB, procHandler, procAdapter, procScheduler)

	// todo запустить дататаски
	processor.RunTasks()
	err := processor.Run()
	if err != nil {
		processor.Stop()
	}
}
