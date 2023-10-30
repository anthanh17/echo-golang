package model

import "time"

type User struct {
	UserId    string
	FullName  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string
}
