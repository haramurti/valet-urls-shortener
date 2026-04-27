package entity

import (
	"time"

	"github.com/google/uuid"
)

type Url struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	OriginalUrl string     `gorm:"type:text;not null" json:"original_url"`
	ShortCode   string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"short_code"`
	UpdatedAt   time.Time  `gorm:"not null" json:"updated_at"`
	CreatedAt   time.Time  `gorm:"not null" json:"created_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
}
