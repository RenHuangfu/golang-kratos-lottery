package repo

import (
	"lottery/common/entity"
)

type ResultRepo interface {
	Get(id uint) (*entity.Result, error)
	GetAll() ([]*entity.Result, error)
	CountAll() (int64, error)
	Create(result *entity.Result) error
	Delete(id uint) error
	DeleteAll() error
	Update(result *entity.Result, cols ...string) error
	GetFromCache(id uint) (*entity.Result, error)
}
