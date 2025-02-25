package entity

import "time"

// BlackIp ip黑明单表
type BlackIp struct {
	Id         uint       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Ip         string     `gorm:"column:ip;type:varchar(50);comment:IP地址;NOT NULL" json:"ip"`
	BlackTime  time.Time  `gorm:"column:black_time;type:datetime;default:1000-01-01 00:00:00;comment:黑名单限制到期时间;NOT NULL" json:"black_time"`
	SysCreated *time.Time `gorm:"autoCreateTime;column:sys_created;type:datetime;default null;comment:创建时间;NOT NULL" json:"sys_created"`
	SysUpdated *time.Time `gorm:"autoUpdateTime;column:sys_updated;type:datetime;default null;comment:修改时间;NOT NULL" json:"sys_updated"`
}

func (m *BlackIp) TableName() string {
	return "t_black_ip"
}
