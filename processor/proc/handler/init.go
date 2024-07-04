package handler

import "frendler/processor/db"

func Init(db db.DB) *HandlerImpl {
	return &HandlerImpl{db}
}
