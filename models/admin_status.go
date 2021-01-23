package models

import "time"

type AdminStatus struct {
	StatusId   int       `xorm:"not null pk autoincr INT(11)"`
	Aid        int       `xorm:"not null default 0 comment('管理员iD') INT(11)"`
	LoginTime  time.Time `xorm:"comment('登录时间') TIMESTAMP"`
	LoginIp    string    `xorm:"not null default '' comment('IP') CHAR(20)"`
	Login      int       `xorm:"not null default 0 comment('登录次数') INT(11)"`
	AidAdd     int       `xorm:"not null default 0 comment('添加人') INT(11)"`
	AidUpdate  int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	TimeUpdate time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Remark     string    `xorm:"not null default '' comment('备注') VARCHAR(255)"`
}