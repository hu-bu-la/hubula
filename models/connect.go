package models

import "time"

type Connect struct {
	ConnectId int       `xorm:"not null pk autoincr INT(11)"`
	TypeId    int       `xorm:"not null default 0 comment('类别id') index INT(11)"`
	Uid       int       `xorm:"not null default 0 comment('用户id') index INT(11)"`
	OpenId    string    `xorm:"not null default '' comment('对应唯一开放id') index CHAR(80)"`
	Token     string    `xorm:"not null default '' comment('开放密钥') VARCHAR(80)"`
	Type      int       `xorm:"not null default 1 comment('登录类型1腾讯QQ2新浪微博') INT(11)"`
	TypeLogin int       `xorm:"not null default 0 comment('登录模块;302前台还是后台301') INT(11)"`
	TimeAdd   time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Extend    string    `xorm:"default '' comment('扩展参数') VARCHAR(5000)"`
}
