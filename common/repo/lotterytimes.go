package repo

import (
	"lottery/common/entity"
)

type LotteryTimesRepo interface {
	Get(id uint) (*entity.LotteryTimes, error)
	GetByUserIDAndDay(uid uint, day uint) (*entity.LotteryTimes, error)
	GetAll() ([]*entity.LotteryTimes, error)
	CountAll() (int64, error)
	Create(lotteryTimes *entity.LotteryTimes) error
	Delete(id uint) error
	DeleteAll() error
	Update(lotteryTimes *entity.LotteryTimes, cols ...string) error
	IncrUserDayLotteryNum(uid uint) int64
	InitUserLuckyNum(uid uint, num int64) error
	ResetIPLotteryNums()
	ResetUserLotteryNums()
}
