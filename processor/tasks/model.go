package tasks

import (
	"fmt"
	"frendler/processor/proc/handler"
	"time"
)

type TaskImpl struct {
	name     string
	interval int
}

func (t TaskImpl) Init(name string) {
}

func (t *TaskImpl) Run(userID int, db handler.HandlerImpl) {
	go func() {
		for {
			contactCount, conflictCount, err := db.DB.GetContactStats(userID)
			if err != nil {
				fmt.Printf("Error getting contact stats: %v\n", err)
			} else {
				fmt.Printf("Contact count: %d, Conflict count: %d\n", contactCount, conflictCount)
			}
			time.Sleep(time.Duration(t.interval) * time.Second)
		}
	}()
}
