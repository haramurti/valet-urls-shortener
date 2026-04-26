package entity

import "time"

type User struct {
	ID             string
	Email          string
	HashedPassword string
	CreatedAt      time.Time
	DeletedAt      *time.Time
}
