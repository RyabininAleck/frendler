package tasks

import (
	"log"
	"time"

	"frendler/processor/db"
)

func CreateContactUpdateTask(db db.DB, interval time.Duration) *ContactUpdateTask {
	return &ContactUpdateTask{interval, db}
}

type ContactUpdateTask struct {
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

				//todo  надо придумать как останавливать. Можно сделать канал для закрытия, можно сделать через контекст. Лучше через контекст. Остановка должна сопровождаться завершением всех работающих транзакций тасок и завершением горутин после их окончания
			}
		}
	}()
}

func (t *ContactUpdateTask) execute() {
	log.Println("Contact Update Task")

}
