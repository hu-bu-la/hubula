package models

import "time"

type Member struct {
	Uid      int       `xorm:"not null pk autoincr INT(11)"`
	Mobile   string    `xorm:"not null default '' index CHAR(11)"`
	Username string    `xorm:"not null default '' comment('用户名') index CHAR(30)"`
	Mail     string    `xorm:"not null default '' comment('邮箱') index CHAR(32)"`
	Password string    `xorm:"not null default '' comment('密码') CHAR(32)"`
	Salt     string    `xorm:"not null default '' comment('干扰码') CHAR(6)"`
	RegIp    string    `xorm:"not null default '' comment('注册IP') CHAR(15)"`
	RegTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('注册时间') TIMESTAMP"`
	IsDel    int       `xorm:"not null default 0 comment('状态0正常1删除') index TINYINT(1)"`
	GroupId  int       `xorm:"not null default 410 comment('用户组ID') index INT(11)"`
	TrueName string    `xorm:"not null default '' comment('真实姓名') VARCHAR(32)"`
	Name     string    `xorm:"not null default '' comment('店铺名称') VARCHAR(100)"`
}
