package tasks

import (
	"frendler/processor/proc/handler"
	"log"
	"time"
)

type TaskImpl struct {
	name     string
	interval int
}

func (t TaskImpl) Init(name string) {
}

func (t *TaskImpl) Run(db *handler.HandlerImpl) {
	go func() {
		for {
			if err := db.GetContactStats; err != nil {
				log.Fatalf("Failed to update friends: %v", err)
			}

			time.Sleep(time.Duration(t.interval) * time.Second)
		}
	}()
}
