package user

import "time"

type User struct {
	ID             string
	Email          string
	HashedPassword string
	createdAt      time.Time
	deletedAt      *time.Time
}
