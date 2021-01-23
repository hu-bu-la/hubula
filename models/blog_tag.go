package models

import "time"

type BlogTag struct {
	TagId   int       `xorm:"not null pk autoincr INT(11)"`
	Name    string    `xorm:"not null default '' comment('名称') CHAR(100)"`
	TimeAdd time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Aid     int       `xorm:"not null default 0 comment('管理员ID') INT(11)"`
	BlogId  int       `xorm:"not null default 0 comment('文章ID') INT(11)"`
}

