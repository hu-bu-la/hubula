package models

import "time"

type Blog struct {
	BlogId      int       `xorm:"not null pk autoincr INT(11)"`
	Aid         int       `xorm:"not null default 0 comment('管理员AID') INT(11)"`
	IsDel       int       `xorm:"not null default 0 comment('是否删除1是0否') index(is_del) TINYINT(1)"`
	IsOpen      int       `xorm:"not null default 1 comment('启用1是0否') index(is_del) TINYINT(1)"`
	Status      int       `xorm:"not null default 0 comment('状态') index(is_del) INT(11)"`
	TimeSystem  time.Time `xorm:"comment('创建时间,系统时间不可修改') TIMESTAMP"`
	TimeUpdate  time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	TimeAdd     time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间,可修改') TIMESTAMP"`
	Title       string    `xorm:"not null default '' comment('标题') VARCHAR(255)"`
	Author      string    `xorm:"not null default '' comment('作者') VARCHAR(255)"`
	Url         string    `xorm:"not null default '' comment('网址') VARCHAR(255)"`
	UrlSource   string    `xorm:"not null default '' comment('来源地址(转载)') VARCHAR(255)"`
	UrlRewrite  string    `xorm:"not null default '' comment('自定义伪静态Url') index CHAR(64)"`
	Description string    `xorm:"not null default '' comment('摘要') VARCHAR(255)"`
	Content     string    `xorm:"not null comment('内容') TEXT"`
	Type        int       `xorm:"not null default 0 comment('类型0文章10001博客栏目') index INT(11)"`
	ModuleId    int       `xorm:"not null default 0 comment('模块10019技术10018生活') index INT(10)"`
	SourceId    int       `xorm:"not null default 0 comment('来源:后台，接口，其他') index INT(11)"`
	TypeId      int       `xorm:"not null default 0 comment('类别ID，原创，转载，翻译') index(is_del) INT(11)"`
	CatId       int       `xorm:"not null default 0 comment('分类ID，栏目') index(is_del) INT(11)"`
	Tag         string    `xorm:"not null default '' comment('标签') VARCHAR(255)"`
	Thumb       string    `xorm:"not null default '' comment('缩略图') VARCHAR(255)"`
	IsRelevant  int       `xorm:"not null default 0 comment('相关文章1是0否') TINYINT(1)"`
	IsJump      int       `xorm:"not null default 0 comment('跳转1是0否') TINYINT(1)"`
	IsComment   int       `xorm:"not null default 1 comment('允许评论1是0否') TINYINT(1)"`
	IsRead      int       `xorm:"not null default 10014 comment('是否阅读10014未看10015在看10016已看') INT(11)"`
	Sort        int       `xorm:"not null default 0 comment('排序') index(is_del) INT(11)"`
	Remark      string    `xorm:"not null default '' comment('备注') VARCHAR(255)"`
}

