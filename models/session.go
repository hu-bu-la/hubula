package models

import "time"

type Session struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Uid        int       `xorm:"not null default 0 comment('用户UID') index(uid) INT(11)"`
	Ip         string    `xorm:"not null default '' comment('IP') CHAR(15)"`
	TimeAdd    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('登录时间') TIMESTAMP"`
	ErrorCount int       `xorm:"not null default 0 comment('密码输入错误次数') TINYINT(1)"`
	AppId      int       `xorm:"not null default 0 comment('登录应用') INT(11)"`
	TypeLogin  int       `xorm:"not null default 0 comment('登录方式;302前台还是后台301') index(uid) INT(11)"`
	Md5        string    `xorm:"not null default '' comment('md5') CHAR(32)"`
	TypeClient int       `xorm:"not null default 0 comment('登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他') index(uid) INT(11)"`
}
