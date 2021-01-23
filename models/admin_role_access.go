package models

type AdminRoleAccess struct {
	Aid    int `xorm:"default 0 comment('管理员ID') unique(aid_role_id) INT(11)"`
	RoleId int `xorm:"default 0 comment('角色ID') unique(aid_role_id) INT(11)"`
}
