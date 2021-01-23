package models

import "time"

type Log struct {
	LogId      int       `xorm:"not null pk autoincr INT(11)"`
	Id         int       `xorm:"not null default 0 comment('id') index INT(11)"`
	Aid        int       `xorm:"not null default 0 comment('管理员ID') index INT(11)"`
	Uid        int       `xorm:"not null default 0 comment('用户id') index INT(11)"`
	TimeAdd    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Mark       string    `xorm:"not null default '' comment('标志自定义标志') CHAR(32)"`
	Data       string    `xorm:"comment('其他内容') TEXT"`
	No         string    `xorm:"not null default '' comment('单号') index CHAR(50)"`
	TypeLogin  int       `xorm:"not null default 0 comment('登录方式;302前台还是后台301') index INT(11)"`
	TypeClient int       `xorm:"not null default 0 comment('登录客户端类别;321电脑;322安卓;323IOS;324手机网页;325其他') index INT(11)"`
	Ip         string    `xorm:"not null default '' comment('IP') CHAR(20)"`
	Msg        string    `xorm:"comment('自定义说明') VARCHAR(255)"`
}
