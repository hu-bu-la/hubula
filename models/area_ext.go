package models

type AreaExt struct {
	ExtId           int    `xorm:"not null pk autoincr INT(11)"`
	Id              int    `xorm:"default 0 comment('ID') index(id) INT(11)"`
	Name            string `xorm:"default '' comment('名称') CHAR(50)"`
	NameEn          string `xorm:"default '' comment('英文名称') VARCHAR(100)"`
	ParentId        int    `xorm:"default 0 comment('上级栏目ID') index(id) INT(11)"`
	Type            int    `xorm:"default 0 comment('类别;0默认;1又名;2;3属于;11已合并到;12已更名为') TINYINT(4)"`
	NameTraditional string `xorm:"default '' comment('繁体名称') VARCHAR(50)"`
	Sort            int    `xorm:"default 0 comment('排序') INT(11)"`
	TypeName        string `xorm:"default '' comment('类别名称') VARCHAR(50)"`
	OtherName       string `xorm:"default '' comment('根据类别名称填写') VARCHAR(50)"`
}

