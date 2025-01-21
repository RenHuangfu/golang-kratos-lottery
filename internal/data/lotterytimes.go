package data

import (
	"context"
	"fmt"
	"github.com/BitofferHub/pkg/middlewares/cache"
	"github.com/BitofferHub/pkg/middlewares/log"
	"gorm.io/gorm"
	"lottery/common/constant"
	"lottery/common/entity"
	"lottery/common/repo"
	"math"
)

type lotteryTimesRepo struct {
	data *Data
}

func NewLotteryTimesRepo(data *Data) repo.LotteryTimesRepo {
	return &lotteryTimesRepo{
		data: data,
	}
}

func (r *lotteryTimesRepo) Get(id uint) (*entity.LotteryTimes, error) {
	db := r.data.db
	lotteryTimes := &entity.LotteryTimes{
		Id: id,
	}
	err := db.Model(&entity.LotteryTimes{}).First(lotteryTimes).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, fmt.Errorf("lotteryTimesRepo|Get:%v", err)
	}
	return lotteryTimes, nil
}

func (r *lotteryTimesRepo) GetByUserIDAndDay(uid uint, day uint) (*entity.LotteryTimes, error) {
	db := r.data.db
	lotteryTimes := &entity.LotteryTimes{}
	err := db.Model(&entity.LotteryTimes{}).Where("user_id=? and day=?", uid, day).First(lotteryTimes).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, fmt.Errorf("lotteryTimesRepo|GetByUserID:%v", err)
	}
	return lotteryTimes, nil
}

func (r *lotteryTimesRepo) GetAll() ([]*entity.LotteryTimes, error) {
	db := r.data.db
	var lotteryTimesList []*entity.LotteryTimes
	err := db.Model(&entity.LotteryTimes{}).Where("").Order("sys_updated desc").Find(&lotteryTimesList).Error
	if err != nil {
		return nil, fmt.Errorf("lotteryTimesRepo|GetAll:%v", err)
	}
	return lotteryTimesList, nil
}

func (r *lotteryTimesRepo) CountAll() (int64, error) {
	db := r.data.db
	var num int64
	err := db.Model(&entity.LotteryTimes{}).Count(&num).Error
	if err != nil {
		return 0, fmt.Errorf("lotteryTimesRepo|CountAll:%v", err)
	}
	return num, nil
}

func (r *lotteryTimesRepo) Create(lotteryTimes *entity.LotteryTimes) error {
	db := r.data.db
	err := db.Model(&entity.LotteryTimes{}).Create(lotteryTimes).Error
	if err != nil {
		return fmt.Errorf("lotteryTimesRepo|Create:%v", err)
	}
	return nil
}

func (r *lotteryTimesRepo) Delete(id uint) error {
	db := r.data.db
	lotteryTimes := &entity.LotteryTimes{Id: id}
	if err := db.Model(&entity.LotteryTimes{}).Delete(lotteryTimes).Error; err != nil {
		return fmt.Errorf("lotteryTimesRepo|Delete:%v", err)
	}
	return nil
}

func (r *lotteryTimesRepo) DeleteAll() error {
	db := r.data.db
	if err := db.Exec("DELETE FROM t_lottery_times").Error; err != nil {
		log.Errorf("lotteryTimesRepo|DeleteAll:%v", err)
		return fmt.Errorf("lotteryTimesRepo|DeleteAll:%v", err)
	}
	return nil
}

func (r *lotteryTimesRepo) Update(lotteryTimes *entity.LotteryTimes, cols ...string) error {
	var err error
	db := r.data.db
	if len(cols) == 0 {
		err = db.Model(lotteryTimes).Updates(lotteryTimes).Error
	} else {
		err = db.Model(lotteryTimes).Select(cols).Updates(lotteryTimes).Error
	}
	if err != nil {
		return fmt.Errorf("lotteryTimesRepo|Update:%v", err)
	}
	return nil
}

// IncrUserDayLotteryNum 每天缓存的用户抽奖次数递增，返回递增后的数值
func (r *lotteryTimesRepo) IncrUserDayLotteryNum(uid uint) int64 {
	redisCli := r.data.cache
	i := uid % constant.UserFrameSize
	// 集群的redis统计数递增
	key := fmt.Sprintf(constant.UserLotteryDayNumPrefix+"%d", i)
	ret, err := redisCli.HIncrBy(context.Background(), key, fmt.Sprint(uid), 1)
	if err != nil {
		log.Errorf("lotteryTimesRepo|IncrUserDayLotteryNum:%v", err)
		return math.MaxInt32
	}
	return ret
}

// InitUserLuckyNum 从给定的数据直接初始化用户的参与抽奖次数
func (r *lotteryTimesRepo) InitUserLuckyNum(uid uint, num int64) error {
	redisCli := r.data.cache
	if num <= 1 {
		return nil
	}
	i := uid % constant.UserFrameSize
	key := fmt.Sprintf(constant.UserLotteryDayNumPrefix+"%d", i)
	_, err := redisCli.HSet(context.Background(), key, fmt.Sprint(uid), num)
	if err != nil {
		log.Errorf("lotteryTimesRepo|InitUserLuckyNum:%v", err)
		return fmt.Errorf("lotteryTimesRepo|InitUserLuckyNum:%v", err)
	}
	return nil
}

func (r *lotteryTimesRepo) ResetIPLotteryNums() {
	//log.Infof("重置所有的IP抽奖次数")
	for i := 0; i < constant.IpFrameSize; i++ {
		key := fmt.Sprintf("day_ip_num_%d", i)
		if err := cache.GetRedisCli().Delete(context.Background(), key); err != nil {
			log.Errorf("ResetIPLotteryNums err:%v", err)
		}
	}
	//log.Infof("重置所有的IP抽奖次数完成！！！")
}

func (r *lotteryTimesRepo) ResetUserLotteryNums() {
	//log.Infof("重置今日用户抽奖次数")
	for i := 0; i < constant.UserFrameSize; i++ {
		key := fmt.Sprintf(constant.UserLotteryDayNumPrefix+"%d", i)
		if err := cache.GetRedisCli().Delete(context.Background(), key); err != nil {
			log.Errorf("ResetIPLotteryNums err:%v", err)
		}
	}
}
