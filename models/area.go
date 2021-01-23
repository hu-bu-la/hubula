package models

type Area struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	Name            string `xorm:"default '' comment('名称') CHAR(50)"`
	NameEn          string `xorm:"default '' comment('英文名称') VARCHAR(100)"`
	ParentId        int    `xorm:"default 0 comment('上级栏目ID') index INT(11)"`
	Type            int    `xorm:"default 0 comment('类别;0默认;') TINYINT(4)"`
	NameTraditional string `xorm:"default '' comment('繁体名称') VARCHAR(50)"`
	Sort            int    `xorm:"default 0 comment('排序') INT(11)"`
}

