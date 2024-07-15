package main

import (
	integrationAdapter "frendler/processor/adapter"
	"frendler/processor/config"
	database "frendler/processor/db"
	"frendler/processor/proc"
	handler "frendler/processor/proc/handler"
	dataTasks "frendler/processor/tasks"
	"frendler/storage"
)

func main() {

	cfg := config.Get()

	db := database.Init(cfg.DB)
	storage.Migrations(db)

	hadler := handler.Init(db)

	// todo подключиться к адаптеру, настроить
	adapter := integrationAdapter.Init(cfg.Adapter)
	// todo подключиться дататаски
	tasks := dataTasks.Init(cfg.Task)

	processor := proc.Init(cfg, db, hadler, adapter, tasks)

	// todo запустить дататаски
	processor.RunTasks()
	err := processor.Run()
	if err != nil {
		processor.Stop()
	}
}
