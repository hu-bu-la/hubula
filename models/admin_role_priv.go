package models

type AdminRolePriv struct {
	Id     int    `xorm:"not null pk autoincr INT(10)"`
	RoleId int    `xorm:"not null default 0 comment('角色ID') index index(role_id_2) SMALLINT(3)"`
	S      string `xorm:"not null default '' comment('模块/控制器/动作') index(role_id_2) CHAR(100)"`
	Data   string `xorm:"not null default '' comment('其他参数') CHAR(50)"`
	Aid    int    `xorm:"not null default 0 comment('管理员ID') INT(10)"`
	MenuId int    `xorm:"not null default 0 comment('菜单ID') INT(10)"`
	Type   string `xorm:"not null default 'url' comment('类别url菜单function独立功能user用户独有') CHAR(32)"`
}

