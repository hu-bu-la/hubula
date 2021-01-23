package models

import "time"

type Template struct {
	TemplateId int       `xorm:"not null pk autoincr comment('模板ID') INT(11)"`
	Name       string    `xorm:"not null default '' comment('模板名称(中文)') VARCHAR(80)"`
	Mark       string    `xorm:"not null default '' comment('模板名称标志(英文)（调用时使用）') VARCHAR(80)"`
	Title      string    `xorm:"not null default '' comment('邮件标题') VARCHAR(255)"`
	Type       int       `xorm:"not null default 0 comment('模板类型1短信模板2邮箱模板') TINYINT(1)"`
	Use        int       `xorm:"not null default 0 comment('用途') TINYINT(2)"`
	Content    string    `xorm:"TEXT"`
	Remark     string    `xorm:"not null default '' comment('备注') VARCHAR(1024)"`
	TimeAdd    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	TimeUpdate time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	CodeNum    int       `xorm:"not null default 0 comment('验证码位数') TINYINT(3)"`
	Aid        int       `xorm:"not null default 0 comment('添加人') INT(11)"`
}
