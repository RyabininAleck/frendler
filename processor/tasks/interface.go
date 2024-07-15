package tasks

import "frendler/processor/proc/handler"

type Task interface {
	Init(string)
	Run(*handler.HandlerImpl)
	//Stop()
}
