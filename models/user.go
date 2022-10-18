package models

import "time"

type User struct {
	ID        int32
	Name      string
	CreatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
