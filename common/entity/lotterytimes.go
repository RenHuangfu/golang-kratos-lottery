package entity

import "time"

// LotteryTimes 用户每日抽奖次数表
type LotteryTimes struct {
	Id         uint       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserId     uint       `gorm:"column:user_id;type:int(10) unsigned;default:0;comment:用户ID;NOT NULL" json:"user_id"`
	Day        uint       `gorm:"column:day;type:int(10) unsigned;default:0;comment:日期，如：20220625;NOT NULL" json:"day"`
	Num        uint       `gorm:"column:num;type:int(10) unsigned;default:0;comment:次数;NOT NULL" json:"num"`
	SysCreated *time.Time `gorm:"autoCreateTime;column:sys_created;type:datetime;default null;comment:创建时间;NOT NULL" json:"sys_created"`
	SysUpdated *time.Time `gorm:"autoUpdateTime;column:sys_updated;type:datetime;default null;comment:修改时间;NOT NULL" json:"sys_updated"`
}

func (l *LotteryTimes) TableName() string {
	return "t_lottery_times"
}
