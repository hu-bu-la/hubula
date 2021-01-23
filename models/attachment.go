package models

import "time"

type Attachment struct {
	AttachmentId int       `xorm:"not null pk autoincr comment('附件ID') INT(10)"`
	Module       string    `xorm:"not null default '' comment('模块') index CHAR(32)"`
	Mark         string    `xorm:"not null default '' comment('标记标志') index CHAR(60)"`
	TypeId       int       `xorm:"not null default 0 comment('类别ID') INT(5)"`
	Name         string    `xorm:"not null default '' comment('保存的文件名称') CHAR(50)"`
	NameOriginal string    `xorm:"not null default '' comment('原文件名') VARCHAR(255)"`
	Path         string    `xorm:"not null default '' comment('文件路径') CHAR(200)"`
	Size         int       `xorm:"not null default 0 comment('文件大小') INT(10)"`
	Ext          string    `xorm:"not null default '' comment('文件后缀') CHAR(10)"`
	IsImage      int       `xorm:"not null default 0 comment('是否图片1是0否') TINYINT(1)"`
	IsThumb      int       `xorm:"not null default 0 comment('是否缩略图1是0否') TINYINT(1)"`
	Downloads    int       `xorm:"not null default 0 comment('下载次数') INT(8)"`
	TimeAdd      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间上传时间') TIMESTAMP"`
	Ip           string    `xorm:"not null default '' comment('上传IP') CHAR(15)"`
	Status       int       `xorm:"not null default 0 comment('状态99正常;') index TINYINT(2)"`
	Md5          string    `xorm:"not null default '' comment('md5') index CHAR(32)"`
	Sha1         string    `xorm:"not null default '' comment('sha1') CHAR(40)"`
	Id           int       `xorm:"not null default 0 comment('所属ID') index INT(10)"`
	Aid          int       `xorm:"not null default 0 comment('后台管理员ID') index INT(10)"`
	Uid          int       `xorm:"not null default 0 comment('前台用户ID') index INT(10)"`
	IsShow       int       `xorm:"not null default 1 comment('是否显示1是0否') index TINYINT(1)"`
	Http         string    `xorm:"not null default '' comment('图片http地址') VARCHAR(100)"`
}

