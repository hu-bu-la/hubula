package models

import "time"

type BlogSyncQueue struct {
	QueueId    int       `xorm:"not null pk autoincr INT(11)"`
	BlogId     int       `xorm:"not null default 0 comment('本站博客id') INT(11)"`
	TypeId     int       `xorm:"not null default 0 comment('类型') INT(11)"`
	Status     int       `xorm:"not null default 0 comment('状态：0:待运行 10:失败 99:成功') TINYINT(3)"`
	TimeUpdate time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('最后一次更新时间') TIMESTAMP"`
	TimeAdd    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('插入时间') TIMESTAMP"`
	Msg        string    `xorm:"not null default '' comment('内容') VARCHAR(255)"`
	MapId      int       `xorm:"not null default 0 comment('同步ID') INT(11)"`
}

