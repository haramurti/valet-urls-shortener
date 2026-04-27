package contract

import (
	"github.com/google/uuid"
	"github.com/haramurti/valeth-urls-shortener/Internal/app/url/entity"
)

type UrlRepository interface {
	Create(url *entity.Url) error
	FindByUrlId(urlId uuid.UUID) (*entity.Url, error)
	Update(url *entity.Url) error
	Delete(url *entity.Url) error
}
