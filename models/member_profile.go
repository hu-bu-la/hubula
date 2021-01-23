package models

type MemberProfile struct {
	ProfileId   int    `xorm:"not null pk autoincr INT(11)"`
	Uid         int    `xorm:"not null default 0 comment('UID') index INT(11)"`
	Sex         int    `xorm:"not null default 1 comment('性别1男2女3中性0保密') TINYINT(1)"`
	Job         string `xorm:"not null default '' comment('担任职务') VARCHAR(50)"`
	Qq          string `xorm:"not null default '' VARCHAR(20)"`
	Phone       string `xorm:"not null default '' comment('电话') VARCHAR(20)"`
	County      int    `xorm:"not null default 1 comment('国家') INT(11)"`
	Province    int    `xorm:"not null default 0 comment('省') INT(11)"`
	City        int    `xorm:"not null default 0 comment('市') INT(11)"`
	District    int    `xorm:"not null default 0 comment('区') INT(11)"`
	Address     string `xorm:"not null default '' comment('地址') VARCHAR(255)"`
	Wechat      string `xorm:"not null default '' comment('微信') VARCHAR(20)"`
	RemarkAdmin string `xorm:"comment('客服备注') TEXT"`
}

