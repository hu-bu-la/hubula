package models

type MemberGroupExt struct {
	GroupId int `xorm:"not null pk autoincr comment('用户ID') INT(10)"`
}

