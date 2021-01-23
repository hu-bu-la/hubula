package models

type AdminMenu struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	Name     string `xorm:"not null default '' comment('名称') CHAR(100)"`
	ParentId int    `xorm:"not null default 0 comment('上级菜单') index INT(11)"`
	S        string `xorm:"not null default '' comment('模块/控制器/动作') index CHAR(60)"`
	Data     string `xorm:"not null default '' comment('其他参数') CHAR(100)"`
	Sort     int    `xorm:"not null default 0 comment('排序') index INT(11)"`
	Remark   string `xorm:"not null default '' comment('备注') VARCHAR(255)"`
	Type     string `xorm:"not null default 'url' comment('类别url菜单function独立功能user用户独有') CHAR(32)"`
	Level    int    `xorm:"not null default 0 comment('级别') TINYINT(3)"`
	Level1Id int    `xorm:"not null default 0 comment('1级栏目ID') INT(11)"`
	Md5      string `xorm:"not null default '' comment('s的md5值') CHAR(32)"`
	IsShow   int    `xorm:"not null default 1 comment('显示隐藏;1显示;0隐藏') TINYINT(1)"`
	IsUnique int    `xorm:"not null default 0 comment('用户独有此功能1是0否') TINYINT(1)"`
}

