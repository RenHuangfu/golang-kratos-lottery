package usecase

import (
	"context"
	"fmt"
	"github.com/BitofferHub/pkg/middlewares/cache"
	"github.com/BitofferHub/pkg/middlewares/log"
	"lottery/common/constant"
	entity2 "lottery/common/entity"
	repo2 "lottery/common/repo"
	"lottery/internal/utils"
	"math"
	"strconv"
	"time"
)

type LimitCase struct {
	lotteryTimesRepo repo2.LotteryTimesRepo
	blackIpRepo      repo2.BlackIpRepo
	blackUserRepo    repo2.BlackUserRepo
}

func NewLimitCase(bur repo2.BlackUserRepo, bir repo2.BlackIpRepo, ltr repo2.LotteryTimesRepo) *LimitCase {
	return &LimitCase{
		blackUserRepo:    bur,
		blackIpRepo:      bir,
		lotteryTimesRepo: ltr,
	}
}

// GetUserCurrentLotteryTimes 获取当天该用户的抽奖次数
func (l *LimitCase) GetUserCurrentLotteryTimes(ctx context.Context, uid uint) (*entity2.LotteryTimes, error) {
	y, m, d := time.Now().Date()
	strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
	day, _ := strconv.Atoi(strDay)
	lotteryTimes, err := l.lotteryTimesRepo.GetByUserIDAndDay(uid, uint(day))
	if err != nil {
		log.ErrorContextf(ctx, "lotteryTimesCase|GetUserCurrentLotteryTimes:%v", err)
		return nil, err
	}
	return lotteryTimes, nil
}

// CheckUserDayLotteryTimes 判断当天是否还可以进行抽奖
func (l *LimitCase) CheckUserDayLotteryTimes(ctx context.Context, uid uint) (bool, error) {
	userLotteryTimes, err := l.GetUserCurrentLotteryTimes(ctx, uid)
	if err != nil {
		return false, fmt.Errorf("checkUserDayLotteryTimes|err:%v", err)
	}
	if userLotteryTimes != nil {
		// 今天的抽奖记录已经达到了抽奖次数限制
		if userLotteryTimes.Num >= constant.UserPrizeMax {
			return false, nil
		} else {
			userLotteryTimes.Num++
			if err := l.lotteryTimesRepo.Update(userLotteryTimes, "num"); err != nil {
				return false, fmt.Errorf("updateLotteryTimes｜update:%v", err)
			}
		}
		return true, nil
	}
	y, m, d := time.Now().Date()
	strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
	day, _ := strconv.Atoi(strDay)
	lotteryTimesInfo := &entity2.LotteryTimes{
		UserId: uid,
		Day:    uint(day),
		Num:    1,
	}
	if err := l.lotteryTimesRepo.Create(lotteryTimesInfo); err != nil {
		return false, fmt.Errorf("updateLotteryTimes｜create:%v", err)
	}
	return true, nil
}

func (l *LimitCase) CheckUserDayLotteryTimesWithCache(ctx context.Context, uid uint) (bool, error) {
	// 通过缓存验证
	userLotteryNum := l.lotteryTimesRepo.IncrUserDayLotteryNum(uid)
	//log.InfoContextf(ctx, "CheckUserDayLotteryTimesWithCache|userLotteryNum = %d", userLotteryNum)
	// 缓存验证没通过，直接返回
	log.Infof("checkUserDayLotteryTimes|uid=%d|userLotteryNum=%d", uid, userLotteryNum)
	if userLotteryNum > constant.UserPrizeMax {
		return false, nil
	}
	// 通过数据库验证，还要在数据库中做一次验证
	userLotteryTimes, err := l.GetUserCurrentLotteryTimes(ctx, uid)
	if err != nil {
		return false, fmt.Errorf("checkUserDayLotteryTimes|err:%v", err)
	}
	if userLotteryTimes != nil {
		// 数据库验证今天的抽奖记录已经达到了抽奖次数限制，不能在抽奖
		if userLotteryTimes.Num >= constant.UserPrizeMax {
			// 缓存数据不可靠，不对，需要更新
			if int64(userLotteryTimes.Num) > userLotteryNum {
				if err = l.lotteryTimesRepo.InitUserLuckyNum(uid, int64(userLotteryTimes.Num)); err != nil {
					return false, fmt.Errorf("LimitCase|CheckUserDayLotteryTimesWithCache:%v", err)
				}
			}
			return false, nil
		} else { // 数据库验证通过，今天还可以抽奖
			userLotteryTimes.Num++
			// 此时次数抽奖次数增加了，需要更新缓存
			if int64(userLotteryTimes.Num) > userLotteryNum {
				if err = l.lotteryTimesRepo.InitUserLuckyNum(uid, int64(userLotteryTimes.Num)); err != nil {
					return false, fmt.Errorf("LimitCase|CheckUserDayLotteryTimesWithCache:%v", err)
				}
			}
			// 更新数据库
			if err = l.lotteryTimesRepo.Update(userLotteryTimes); err != nil {
				return false, fmt.Errorf("updateLotteryTimes｜update:%v", err)
			}
		}
		return true, nil
	}
	y, m, d := time.Now().Date()
	strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
	day, _ := strconv.Atoi(strDay)
	lotteryTimesInfo := &entity2.LotteryTimes{
		UserId: uid,
		Day:    uint(day),
		Num:    1,
	}
	if err = l.lotteryTimesRepo.Create(lotteryTimesInfo); err != nil {
		return false, fmt.Errorf("updateLotteryTimes｜create:%v", err)
	}
	if err = l.lotteryTimesRepo.InitUserLuckyNum(uid, 1); err != nil {
		return false, fmt.Errorf("LimitCase|CheckUserDayLotteryTimesWithCache:%v", err)
	}
	return true, nil
}

// CheckIPLimit 验证ip抽奖是否受限制
func (l *LimitCase) CheckIPLimit(ctx context.Context, strIp string) int64 {
	ip := utils.Ip4toInt(strIp)
	i := ip % constant.IpFrameSize
	key := fmt.Sprintf("day_ip_num_%d", i)
	ret, err := cache.GetRedisCli().HIncrBy(ctx, key, strIp, 1)
	if err != nil {
		log.ErrorContextf(ctx, "CheckIPLimit|Incr:%v", err)
		return math.MaxInt32
	}
	return ret
}

func (l *LimitCase) CheckBlackIP(ctx context.Context, ip string) (bool, *entity2.BlackIp, error) {
	info, err := l.blackIpRepo.GetByIP(ip)
	if err != nil {
		log.ErrorContextf(ctx, "CheckBlackIP|GetByIP:%v", err)
		return false, nil, fmt.Errorf("CheckBlackIP|GetByIP:%v", err)
	}
	if info == nil || info.Ip == "" {
		return true, nil, nil
	}
	if time.Now().Before(info.BlackTime) {
		// IP黑名单存在，而且还在黑名单有效期内
		return false, info, nil
	}
	return true, info, nil
}

func (l *LimitCase) CheckBlackIPWithCache(ctx context.Context, ip string) (bool, *entity2.BlackIp, error) {
	info, err := l.blackIpRepo.GetByIPWithCache(ip)
	if err != nil {
		log.ErrorContextf(ctx, "CheckBlackIP|GetByIP:%v", err)
		return false, nil, fmt.Errorf("CheckBlackIP|GetByIP:%v", err)
	}
	if info == nil || info.Ip == "" {
		return true, nil, nil
	}
	if time.Now().Before(info.BlackTime) {
		// IP黑名单存在，而且还在黑名单有效期内
		return false, info, nil
	}
	return true, info, nil
}

func (l *LimitCase) CheckBlackUser(ctx context.Context, uid uint) (bool, *entity2.BlackUser, error) {
	info, err := l.blackUserRepo.GetByUserID(uid)
	if err != nil {
		log.ErrorContextf(ctx, "CheckBlackUser|Get:%v", err)
		return false, nil, fmt.Errorf("CheckBlackUser|Get:%v", err)
	}
	// 黑名单存在并且有效，不能通过
	if info != nil && time.Now().Before(info.BlackTime) {
		return false, info, nil
	}
	return true, info, nil
}

func (l *LimitCase) CheckBlackUserWithCache(ctx context.Context, uid uint) (bool, *entity2.BlackUser, error) {
	info, err := l.blackUserRepo.GetByUserIDWithCache(uid)
	if err != nil {
		log.ErrorContextf(ctx, "CheckBlackUser|Get:%v", err)
		return false, nil, fmt.Errorf("CheckBlackUser|Get:%v", err)
	}
	if info != nil && info.BlackTime.Unix() > time.Now().Unix() {
		// 黑名单存在并且有效
		return false, info, nil
	}
	return true, info, nil
}

func (l *LimitCase) CronJobResetIPLotteryNums() {
	l.lotteryTimesRepo.ResetIPLotteryNums()
}

func (l *LimitCase) CronJobResetUserLotteryNums() {
	l.lotteryTimesRepo.ResetUserLotteryNums()
}
