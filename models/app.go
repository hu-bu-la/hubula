package models

import "time"

type App struct {
	AppId   int       `xorm:"not null pk autoincr INT(11)"`
	TypeId  int       `xorm:"not null default 0 comment('app_id,来源type表') unique INT(11)"`
	Name    string    `xorm:"not null default '' comment('名称') VARCHAR(100)"`
	Mark    string    `xorm:"not null default '' comment('标志') CHAR(32)"`
	Setting string    `xorm:"comment('扩展参数') VARCHAR(5000)"`
	Remark  string    `xorm:"comment('备注') VARCHAR(255)"`
	TimeAdd time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	IsDel   int       `xorm:"not null default 0 comment('是否删除0否1是') INT(11)"`
}

