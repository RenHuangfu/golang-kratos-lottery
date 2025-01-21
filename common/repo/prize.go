package repo

import (
	"lottery/common/entity"
)

type PrizeRepo interface {
	Get(id uint) (*entity.Prize, error)
	GetWithCache(id uint) (*entity.Prize, error)
	GetAll() ([]*entity.Prize, error)
	GetAllWithCache() ([]*entity.Prize, error)
	CountAll() (int64, error)
	CountAllWithCache() (int64, error)
	Create(prize *entity.Prize) error
	CreateInBatches(prizeList []entity.Prize) error
	CreateWithCache(prize *entity.Prize) error
	Delete(id uint) error
	DeleteAll() error
	DeleteWithCache(id uint) error
	Update(prize *entity.Prize, cols ...string) error
	UpdateWithCache(prize *entity.Prize, cols ...string) error
	GetFromCache(id uint) (*entity.Prize, error)
	GetAllUsefulPrizeList() ([]*entity.Prize, error)
	GetAllUsefulPrizeListWithCache() ([]*entity.Prize, error)
	DecrLeftNum(id int, num int) (bool, error)
	DecrLeftNumByPool(prizeID int) (int64, error)
	IncrLeftNum(id int, column string, num int) error
	SetAllByCache(prizeList []*entity.Prize) error
	GetAllByCache() ([]*entity.Prize, error)
	UpdateByCache(prize *entity.Prize) error
	GetPrizePoolNum(prizeID uint) (int, error)
	SetPrizePoolNum(key string, prizeID uint, num int) error
	IncrPrizePoolNum(key string, prizeID uint, num int) (int, error)
}
