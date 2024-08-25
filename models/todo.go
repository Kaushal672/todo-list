package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	TodoId    int          `json:"todoId"`
	Title     string       `json:"title"     validate:"required,min=5,max=50"`
	Status    string       `json:"status"    validate:"required,oneof=not_started in_progress completed"`
	UserId    int          `json:"userId"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}
