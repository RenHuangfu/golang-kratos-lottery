package data

import (
	"github.com/BitofferHub/pkg/middlewares/cache"
	"github.com/BitofferHub/pkg/middlewares/gormcli"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"lottery/internal/conf"
)

type Data struct {
	db    *gorm.DB
	cache *cache.Client
}

func NewData(db *gorm.DB, cache *cache.Client) *Data {
	dt := &Data{db: db, cache: cache}
	return dt
}

func NewDatabase(conf *conf.Data) *gorm.DB {
	dt := conf.GetDatabase()
	gormcli.Init(
		gormcli.WithAddr(dt.GetAddr()),
		gormcli.WithUser(dt.GetUser()),
		gormcli.WithPassword(dt.GetPassword()),
		gormcli.WithDataBase(dt.GetDatabase()),
		gormcli.WithMaxIdleConn(int(dt.GetMaxIdleConn())),
		gormcli.WithMaxOpenConn(int(dt.GetMaxOpenConn())),
		gormcli.WithMaxIdleTime(int64(dt.GetMaxIdleTime())),
		// 如果设置了慢查询阈值，就打印日志
		gormcli.WithSlowThresholdMillisecond(dt.GetSlowThresholdMillisecond()),
	)

	return gormcli.GetDB()
}

func NewCache(conf *conf.Data) *cache.Client {
	dt := conf.GetRedis()
	cache.Init(
		cache.WithAddr(dt.GetAddr()),
		cache.WithPassWord(dt.GetPassword()),
		cache.WithDB(int(dt.GetDb())),
		cache.WithPoolSize(int(dt.GetPoolSize())))

	return cache.GetRedisCli()
}
