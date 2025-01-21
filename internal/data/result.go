package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BitofferHub/pkg/middlewares/log"
	"gorm.io/gorm"
	"lottery/common/entity"
	"lottery/common/repo"
	"strconv"
)

type resultRepo struct {
	data *Data
}

func NewResultRepo(data *Data) repo.ResultRepo {
	return &resultRepo{
		data: data,
	}
}

func (r *resultRepo) Get(id uint) (*entity.Result, error) {
	db := r.data.db
	// 优先从缓存获取
	result, err := r.GetFromCache(id)
	if err == nil && result != nil {
		return result, nil
	}
	result = &entity.Result{
		Id: id,
	}
	err = db.Model(&entity.Result{}).First(result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, fmt.Errorf("resultRepo|Get:%v", err)
	}
	return result, nil
}

func (r *resultRepo) GetAll() ([]*entity.Result, error) {
	db := r.data.db
	var results []*entity.Result
	err := db.Model(&entity.Result{}).Where("").Order("sys_updated desc").Find(&results).Error
	if err != nil {
		return nil, fmt.Errorf("resultRepo|GetAll:%v", err)
	}
	return results, nil
}

func (r *resultRepo) CountAll() (int64, error) {
	db := r.data.db
	var num int64
	err := db.Model(&entity.Result{}).Count(&num).Error
	if err != nil {
		return 0, fmt.Errorf("resultRepo|CountAll:%v", err)
	}
	return num, nil
}

func (r *resultRepo) Create(result *entity.Result) error {
	db := r.data.db
	err := db.Model(&entity.Result{}).Create(result).Error
	if err != nil {
		return fmt.Errorf("resultRepo|Create:%v", err)
	}
	return nil
}

func (r *resultRepo) Delete(id uint) error {
	db := r.data.db
	result := &entity.Result{Id: id}
	if err := db.Model(&entity.Result{}).Delete(result).Error; err != nil {
		return fmt.Errorf("resultRepo|Delete:%v")
	}
	return nil
}

func (r *resultRepo) DeleteAll() error {
	db := r.data.db
	if err := db.Exec("DELETE FROM t_result").Error; err != nil {
		log.Errorf("resultRepo|DeleteAll:%v", err)
		return fmt.Errorf("resultRepo|DeleteAll:%v", err)
	}
	return nil
}

func (r *resultRepo) Update(result *entity.Result, cols ...string) error {
	db := r.data.db
	var err error
	if len(cols) == 0 {
		err = db.Model(result).Updates(result).Error
	} else {
		err = db.Model(result).Select(cols).Updates(result).Error
	}
	if err != nil {
		return fmt.Errorf("resultRepo|Update:%v", err)
	}
	return nil
}

// GetFromCache 根据id从缓存获取奖品
func (r *resultRepo) GetFromCache(id uint) (*entity.Result, error) {
	redisCli := r.data.cache
	idStr := strconv.FormatUint(uint64(id), 10)
	ret, exist, err := redisCli.Get(context.Background(), idStr)
	if err != nil {
		log.Errorf("resultRepo|GetFromCache:" + err.Error())
		return nil, err
	}

	if !exist {
		return nil, nil
	}

	result := entity.Result{}
	json.Unmarshal([]byte(ret), &entity.Result{})

	return &result, nil
}
