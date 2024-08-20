package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserId    int          `json:"userId"`
	Name      string       `json:"name"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
}
