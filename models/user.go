package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserId    int          `json:"userId"`
	Name      string       `json:"name"      binding:"required,min=5,max=16"`
	Password  string       `json:"password"  binding:"required,password"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}
