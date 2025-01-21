package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BitofferHub/pkg/middlewares/log"
	"gorm.io/gorm"
	"lottery/common/constant"
	"lottery/common/entity"
	"lottery/common/repo"
	"lottery/internal/utils"
	"strconv"
)

type blackUserRepo struct {
	data *Data
}

func NewBlackUserRepo(data *Data) repo.BlackUserRepo {
	return &blackUserRepo{
		data: data,
	}
}

func (r *blackUserRepo) GetByUserID(uid uint) (*entity.BlackUser, error) {
	db := r.data.db
	blackUser := &entity.BlackUser{
		UserId: uid,
	}
	err := db.Model(&entity.BlackUser{}).First(blackUser).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, fmt.Errorf("blackUserRepo|Get:%v", err)
	}
	return blackUser, nil
}

func (r *blackUserRepo) GetByUserIDWithCache(uid uint) (*entity.BlackUser, error) {
	// 优先从缓存获取
	blackUser, err := r.GetByCache(uid)
	// 从缓存获取到用户
	if err == nil && blackUser != nil {
		return blackUser, nil
	}
	// 缓存没有获取到黑明单用户
	blackUser = &entity.BlackUser{
		UserId: uid,
	}
	db := r.data.db
	err = db.Model(&entity.BlackUser{}).First(blackUser).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, fmt.Errorf("blackUserRepo|Get:%v", err)
	}
	// db获取到了黑明单用户，同步到缓存中
	if err = r.SetByCache(blackUser); err != nil {
		return nil, fmt.Errorf("blackUserRepo|SetByCache:%v", err)
	}
	return blackUser, nil
}

func (r *blackUserRepo) GetAll() ([]*entity.BlackUser, error) {
	db := r.data.db
	var BlackUsers []*entity.BlackUser
	err := db.Model(&entity.BlackUser{}).Where("").Order("sys_updated desc").Find(&BlackUsers).Error
	if err != nil {
		return nil, fmt.Errorf("blackUserRepo|GetAll:%v", err)
	}
	return BlackUsers, nil
}

func (r *blackUserRepo) CountAll() (int64, error) {
	db := r.data.db
	var num int64
	err := db.Model(&entity.BlackUser{}).Count(&num).Error
	if err != nil {
		return 0, fmt.Errorf("blackUserRepo|CountAll:%v", err)
	}
	return num, nil
}

func (r *blackUserRepo) Create(blackUser *entity.BlackUser) error {
	db := r.data.db
	err := db.Model(blackUser).Create(blackUser).Error
	if err != nil {
		return fmt.Errorf("blackUserRepo|Create:%v", err)
	}
	return nil
}

func (r *blackUserRepo) Delete(id uint) error {
	db := r.data.db
	blackUser := &entity.BlackUser{Id: id}
	if err := db.Model(blackUser).Delete(blackUser).Error; err != nil {
		return fmt.Errorf("blackUserRepo|Delete:%v", err)
	}
	return nil
}

func (r *blackUserRepo) DeleteWithCache(uid uint) error {
	db := r.data.db
	blackUser := &entity.BlackUser{UserId: uid}
	if err := r.UpdateByCache(blackUser); err != nil {
		return fmt.Errorf("blackUserRepo|DeleteWithCache:%v", err)
	}
	if err := db.Model(&entity.BlackUser{}).Delete(blackUser).Error; err != nil {
		return fmt.Errorf("blackUserRepo|Delete:%v", err)
	}
	return nil
}

func (r *blackUserRepo) Update(userID uint, blackUser *entity.BlackUser, cols ...string) error {
	db := r.data.db
	var err error
	if len(cols) == 0 {
		err = db.Model(blackUser).Where("user_id=?", userID).Updates(blackUser).Error
	} else {
		err = db.Model(blackUser).Where("user_id=?", userID).Select(cols).Updates(blackUser).Error
	}
	if err != nil {
		return fmt.Errorf("blackUserRepo|Update:%v", err)
	}
	return nil
}

func (r *blackUserRepo) UpdateWithCache(userID uint, blackUser *entity.BlackUser, cols ...string) error {
	db := r.data.db
	if err := r.UpdateByCache(&entity.BlackUser{UserId: userID}); err != nil {
		return fmt.Errorf("blackUserRepo|DeleteWithCache:%v", err)
	}
	var err error
	if len(cols) == 0 {
		err = db.Model(blackUser).Where("user_id=?", userID).Updates(blackUser).Error
	} else {
		err = db.Model(blackUser).Where("user_id=?", userID).Select(cols).Updates(blackUser).Error
	}
	if err != nil {
		return fmt.Errorf("blackUserRepo|Update:%v", err)
	}
	return nil
}

// GetFromCache 根据id从缓存获取奖品
func (r *blackUserRepo) GetFromCache(id uint) (*entity.BlackUser, error) {
	redisCli := r.data.cache
	idStr := strconv.FormatUint(uint64(id), 10)
	ret, exist, err := redisCli.Get(context.Background(), idStr)
	if err != nil {
		log.Errorf("blackUserRepo|GetFromCache:" + err.Error())
		return nil, err
	}

	if !exist {
		return nil, nil
	}

	blackUser := entity.BlackUser{}
	if err = json.Unmarshal([]byte(ret), &entity.BlackUser{}); err != nil {
		return nil, fmt.Errorf("blackUserRepo|GetFromCache|json.Unmarshal:%v", err)
	}
	return &blackUser, nil
}

func (r *blackUserRepo) GetByCache(uid uint) (*entity.BlackUser, error) {
	redisCli := r.data.cache
	key := fmt.Sprintf(constant.UserCacheKeyPrefix+"%d", uid)
	valueMap, err := redisCli.HGetAll(context.Background(), key)
	if err != nil {
		return nil, fmt.Errorf("blackUserRepo|GetByCache:%v", err)
	}
	userIdStr := valueMap["UserId"]
	num, _ := strconv.Atoi(userIdStr)
	userID := uint(num)
	if userID <= 0 {
		return nil, nil
	}
	idStr := valueMap["Id"]
	id, _ := strconv.Atoi(idStr)
	blackUser := &entity.BlackUser{
		Id:       uint(id),
		UserId:   userID,
		UserName: valueMap["UserName"],
		RealName: valueMap["RealName"],
		Mobile:   valueMap["Mobile"],
		Address:  valueMap["Address"],
		SysIp:    valueMap["SysIp"],
	}
	blackTime, err := utils.ParseTime(valueMap["BlackTime"])
	if err != nil {
		return nil, fmt.Errorf("blackUserRepo|GetByCache:%v", err)
	}
	blackUser.BlackTime = blackTime
	sysCreated, err := utils.ParseTime(valueMap["SysCreated"])
	if err != nil {
		return nil, fmt.Errorf("blackUserRepo|GetByCache:%v", err)
	}
	blackUser.SysCreated = &sysCreated
	sysUpdated, err := utils.ParseTime(valueMap["SysUpdated"])
	if err != nil {
		return nil, fmt.Errorf("blackUserRepo|GetByCache:%v", err)
	}
	blackUser.SysUpdated = &sysUpdated
	return blackUser, nil
}

func (r *blackUserRepo) SetByCache(blackUser *entity.BlackUser) error {
	redisCli := r.data.cache
	if blackUser == nil || blackUser.UserId <= 0 {
		return fmt.Errorf("blackUserRepo|SetByCache invalid user")
	}
	key := fmt.Sprintf(constant.UserCacheKeyPrefix+"%d", blackUser.UserId)
	valueMap := make(map[string]string)
	valueMap["Id"] = strconv.Itoa(int(blackUser.Id))
	valueMap["UserId"] = strconv.Itoa(int(blackUser.UserId))
	valueMap["UserName"] = blackUser.UserName
	valueMap["BlackTime"] = utils.FormatFromUnixTime(blackUser.BlackTime.Unix())
	valueMap["RealName"] = blackUser.RealName
	valueMap["Mobile"] = blackUser.Mobile
	valueMap["Address"] = blackUser.Address
	valueMap["SysCreated"] = utils.FormatFromUnixTime(blackUser.SysCreated.Unix())
	valueMap["SysUpdated"] = utils.FormatFromUnixTime(blackUser.SysUpdated.Unix())
	valueMap["SysIp"] = blackUser.SysIp
	_, err := redisCli.HMSet(context.Background(), key, valueMap)
	if err != nil {
		log.Errorf("blackUserRepo|SetByCache invalid user")
	}
	return nil
}

func (r *blackUserRepo) UpdateByCache(blackUser *entity.BlackUser) error {
	redisCli := r.data.cache
	if blackUser == nil || blackUser.UserId <= 0 {
		return fmt.Errorf("blackUserRepo|UpdateByCache invalid blackUser")
	}
	key := fmt.Sprintf(constant.UserCacheKeyPrefix+"%d", blackUser.UserId)
	if err := redisCli.Delete(context.Background(), key); err != nil {
		return fmt.Errorf("blackUserRepo|UpdateByCache:%v", err)
	}
	return nil
}
