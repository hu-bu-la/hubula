package models

import "time"

type Tag struct {
	TagId   int       `xorm:"not null pk autoincr INT(11)"`
	Name    string    `xorm:"not null default '' comment('名称') CHAR(50)"`
	TimeAdd time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
}
