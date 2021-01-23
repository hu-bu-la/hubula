package models

import "time"

type MemberStatus struct {
	StatusId       int       `xorm:"not null pk autoincr INT(11)"`
	Uid            int       `xorm:"not null default 0 comment('UID') index INT(11)"`
	RegIp          string    `xorm:"not null default '' comment('注册IP') CHAR(15)"`
	RegTime        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('注册时间') TIMESTAMP"`
	RegType        int       `xorm:"not null default 0 comment('注册方式') INT(11)"`
	RegAppId       int       `xorm:"not null default 1 comment('注册来源') INT(11)"`
	LastLoginIp    string    `xorm:"not null default '' comment('最后登录IP') CHAR(15)"`
	LastLoginTime  time.Time `xorm:"comment('最后登录时间') TIMESTAMP"`
	LastLoginAppId int       `xorm:"not null default 0 comment('最后登录app_id') INT(11)"`
	Login          int       `xorm:"not null default 0 comment('登录次数') SMALLINT(5)"`
	IsMobile       int       `xorm:"not null default 0 comment('手机号是否已验证1已验证0未验证') TINYINT(1)"`
	IsEmail        int       `xorm:"not null default 0 comment('邮箱是否已验证1已验证0未验证') TINYINT(1)"`
	AidAdd         int       `xorm:"not null default 0 comment('客服AID') INT(11)"`
}
