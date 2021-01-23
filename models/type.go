package models

import "time"

type Type struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Name      string    `xorm:"not null default '' comment('名称') CHAR(100)"`
	Code      string    `xorm:"not null default '' comment('代码') CHAR(32)"`
	Mark      string    `xorm:"not null default '' comment('标志') index CHAR(32)"`
	TypeId    int       `xorm:"not null default 0 comment('所属类别ID') index INT(11)"`
	ParentId  int       `xorm:"not null default 0 comment('上级ID、属于/上级ID') index INT(11)"`
	Value     int       `xorm:"not null default 0 comment('值') INT(10)"`
	Content   string    `xorm:"not null default '' comment('内容') VARCHAR(255)"`
	IsDel     int       `xorm:"not null default 0 comment('是否删除0否1是') index INT(11)"`
	Sort      int       `xorm:"not null default 0 comment('排序') index INT(11)"`
	Remark    string    `xorm:"comment('备注') VARCHAR(255)"`
	TimeAdd   time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Aid       int       `xorm:"not null default 0 comment('添加人') INT(11)"`
	Module    string    `xorm:"not null default '' comment('模块') CHAR(50)"`
	IsDefault int       `xorm:"not null default 0 comment('是否默认') TINYINT(1)"`
	Setting   string    `xorm:"comment('扩展参数') VARCHAR(255)"`
	IsChild   int       `xorm:"not null default 0 comment('是否有子类1是0否') TINYINT(1)"`
	IsSystem  int       `xorm:"not null default 0 comment('系统参数禁止修改') TINYINT(1)"`
	IsShow    int       `xorm:"not null default 0 comment('是否显示在配置页面上') TINYINT(1)"`
}
