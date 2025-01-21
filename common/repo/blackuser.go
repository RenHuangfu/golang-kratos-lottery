package repo

import (
	"lottery/common/entity"
)

type BlackUserRepo interface {
	GetByUserID(uid uint) (*entity.BlackUser, error)
	GetByUserIDWithCache(uid uint) (*entity.BlackUser, error)
	GetAll() ([]*entity.BlackUser, error)
	CountAll() (int64, error)
	Create(blackUser *entity.BlackUser) error
	Delete(id uint) error
	DeleteWithCache(uid uint) error
	Update(userID uint, blackUser *entity.BlackUser, cols ...string) error
	UpdateWithCache(userID uint, blackUser *entity.BlackUser, cols ...string) error
	GetFromCache(id uint) (*entity.BlackUser, error)
	GetByCache(uid uint) (*entity.BlackUser, error)
	SetByCache(blackUser *entity.BlackUser) error
	UpdateByCache(blackUser *entity.BlackUser) error
}
