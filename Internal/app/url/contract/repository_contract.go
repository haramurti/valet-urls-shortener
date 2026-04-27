package contract

import (
	"github.com/google/uuid"
	"github.com/haramurti/valeth-urls-shortener/Internal/app/url/entity"
)

type UrlRepository interface {
	Create(url *entity.Url) error
	FindByShortCode(shortCode string) (*entity.Url, error)
	FindByUserID(userID uuid.UUID) ([]*entity.Url, error)
	Delete(urlID uuid.UUID) error
}
