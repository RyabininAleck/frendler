package tasks

import (
	"context"
	"log"
	"time"

	"frendler/processor/db"
)

func CreateContactUpdateTask(ctx context.Context, interval time.Duration, db db.DB) *ContactUpdateTask {
	return &ContactUpdateTask{ctx, interval, db}
}

type ContactUpdateTask struct {
	ctx      context.Context
	interval time.Duration
	db       db.DB
}

func (t *ContactUpdateTask) Run() {
	go func() {
		ticker := time.NewTicker(t.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				t.execute()
			case <-t.ctx.Done():
				log.Printf("Stopping Contact update task")
				return
			}
		}
	}()
}

func (t *ContactUpdateTask) execute() {
	log.Println("Contact Update Task")
	userID := 1

	contact, conflictContact, err := t.db.GetContactStats(userID)
	if err != nil {
		log.Fatalf("Failed to updated contact: %v", err)
	} else {
		log.Printf("Contact count: %d, Conflict count: %d\n", contact, conflictContact)
	}
}
