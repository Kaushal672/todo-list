package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	TodoId    int          `json:"todoId"`
	Title     string       `json:"title"`
	Status    string       `json:"status"`
	UserId    int          `json:"userId"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}
