package user

import (
	"time"
)

type (
	User struct {
		ID        string     `json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		Status    string     `json:"status"`
		Password  string     `json:"password"`
	}
)
