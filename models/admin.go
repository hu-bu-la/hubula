package models

import "time"

type Admin struct {
	Aid        int       `xorm:"not null pk autoincr INT(11)"`
	Username   string    `xorm:"not null default '' comment('用户名') index CHAR(30)"`
	Password   string    `xorm:"not null default '' comment('密码') CHAR(32)"`
	Mail       string    `xorm:"not null default '' comment('邮箱') VARCHAR(80)"`
	Salt       string    `xorm:"not null default '' comment('干扰码') VARCHAR(10)"`
	TimeAdd    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('注册时间') TIMESTAMP"`
	TimeUpdate time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Ip         string    `xorm:"not null default '' comment('添加IP') CHAR(15)"`
	JobNo      string    `xorm:"not null default '' comment('工号') VARCHAR(15)"`
	NickName   string    `xorm:"not null default '' comment('昵称') VARCHAR(50)"`
	TrueName   string    `xorm:"not null default '' comment('真实姓名') VARCHAR(50)"`
	Qq         string    `xorm:"not null default '' comment('qq') VARCHAR(50)"`
	Phone      string    `xorm:"not null default '' comment('电话') VARCHAR(50)"`
	Mobile     string    `xorm:"not null default '' comment('手机') VARCHAR(20)"`
	IsDel      int       `xorm:"not null default 0 comment('删除0否1是') index TINYINT(1)"`
}