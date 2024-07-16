package tasks

import (
	"context"
	"log"
	"time"

	"frendler/processor/db"
)

func CreateContactUpdateTask(ctx context.Context, cancelFunc context.CancelFunc, interval time.Duration, db db.DB) *ContactUpdateTask {
	return &ContactUpdateTask{ctx, cancelFunc, interval, db}
}

type ContactUpdateTask struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	interval   time.Duration
	db         db.DB
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

func (t *ContactUpdateTask) Stop() {
	t.cancelFunc()
}

func (t *ContactUpdateTask) execute() {
	userID := 1
	//todo  []socialProfiles(userId, platform,  token) := получить список социальных профилей у которых поле last_contact_updated_at имеет дату раньше чем NOW()-config.platform.updateInterval (с момента последнего обновления прошло больше чем config.platform.updateInterval)    (userId, )
	// for profile(userId, platform) := range  []socialProfiles(userId, platform){
	// пока platform  только google
	// вот тут надо сделать switch platform: case case case  чтобы разделить функции для обновления
	// внутри case google:
	// token, err := GetGoogleToken(socialProfile)
	// client := ...
	// getGoogleContacts(userId, client) из handler/contactUpdate
	// вот тут надо вставить в базу данных новые контакты или обновить те, которые были изменены
	// }
	contact, conflictContact, err := t.db.GetContactStats(userID)
	if err != nil {
		log.Fatalf("Failed to updated contact: %v", err)
	} else {
		log.Printf("Contact count: %d, Conflict count: %d\n", contact, conflictContact)
	}
}
