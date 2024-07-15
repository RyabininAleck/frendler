package tasks

import "frendler/processor/config"

func Init(cfg config.TaskConf) *TaskImpl {
	return &TaskImpl{interval: cfg.Interval}
}
