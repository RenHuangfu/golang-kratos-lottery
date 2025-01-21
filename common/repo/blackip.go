package repo

import (
	"lottery/common/entity"
)

type BlackIpRepo interface {
	Get(id uint) (*entity.BlackIp, error)
	GetByIP(ip string) (*entity.BlackIp, error)
	GetByIPWithCache(ip string) (*entity.BlackIp, error)
	GetAll() ([]*entity.BlackIp, error)
	CountAll() (int64, error)
	Create(blackIp *entity.BlackIp) error
	Delete(id uint) error
	Update(ip string, blackIp *entity.BlackIp, cols ...string) error
	UpdateWithCache(ip string, blackIp *entity.BlackIp, cols ...string) error
	GetFromCache(id uint) (*entity.BlackIp, error)
	SetByCache(blackIp *entity.BlackIp) error
	GetByCache(ip string) (*entity.BlackIp, error)
	UpdateByCache(blackIp *entity.BlackIp) error
}
