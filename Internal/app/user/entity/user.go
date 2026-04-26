package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Email          string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	HashedPassword string     `gorm:"not null" json:"hashed_password"`
	CreatedAt      time.Time  `gorm:"not null" json:"created_at"`
	DeletedAt      *time.Time `gorm:"index" json:"deleted_at"`
}
