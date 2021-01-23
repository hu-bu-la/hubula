package models

import "time"

type BlogSyncMapping struct {
	MapId      int       `xorm:"not null pk autoincr INT(11)"`
	BlogId     int       `xorm:"not null default 0 comment('本站blog的id') INT(11)"`
	TypeId     int       `xorm:"not null default 0 comment('类别id') INT(11)"`
	Id         string    `xorm:"not null default '' comment('csdn的id') VARCHAR(64)"`
	TimeUpdate time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('最后一次更新时间') TIMESTAMP"`
	TimeAdd    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('插入时间') TIMESTAMP"`
	Mark       string    `xorm:"not null default '' comment('标志') CHAR(32)"`
	IsSync     int       `xorm:"not null default 0 comment('是否同步过') TINYINT(1)"`
	Extend     string    `xorm:"comment('扩展参数') VARCHAR(5000)"`
}