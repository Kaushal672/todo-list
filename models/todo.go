package models

import (
	"time"
)

type Todo struct {
	TodoId    int        `json:"todoId"`
	Title     string     `json:"title"               binding:"required,min=5,max=50"`
	Status    string     `json:"status"              binding:"required,oneof=not_started in_progress completed"`
	UserId    int        `json:"userId"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
