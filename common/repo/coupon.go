package repo

import (
	"lottery/common/entity"
)

type CouponRepo interface {
	Get(id uint) (*entity.Coupon, error)
	GetAll() ([]*entity.Coupon, error)
	GetCouponListByPrizeID(prizeID uint) ([]*entity.Coupon, error)
	CountAll() (int64, error)
	Create(coupon *entity.Coupon) error
	Delete(id uint) error
	DeleteAllWithCache() error
	Update(coupon *entity.Coupon, cols ...string) error
	UpdateByCode(code string, coupon *entity.Coupon, cols ...string) error
	GetFromCache(id uint) (*entity.Coupon, error)
	GetGetNextUsefulCoupon(prizeID, couponID int) (*entity.Coupon, error)
	ImportCacheCoupon(prizeID uint, code string) (bool, error)
	ReSetCacheCoupon(prizeID uint) (int64, int64, error)
	GetCacheCouponNum(prizeID uint) (int64, int64, error)
	GetNextUsefulCouponFromCache(prizeID int) (string, error)
}
