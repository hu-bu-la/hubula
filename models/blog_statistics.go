package models

type BlogStatistics struct {
	StatisticsId   int    `xorm:"not null pk autoincr INT(11)"`
	BlogId         int    `xorm:"not null default 0 comment('文章ID') index INT(11)"`
	Comment        int    `xorm:"not null default 0 comment('评论人数') INT(11)"`
	Read           int    `xorm:"not null default 0 comment('阅读人数') INT(11)"`
	SeoTitle       string `xorm:"not null default '' comment('SEO标题') VARCHAR(255)"`
	SeoDescription string `xorm:"not null default '' comment('SEO摘要') VARCHAR(255)"`
	SeoKeyword     string `xorm:"not null default '' comment('SEO关键词') VARCHAR(255)"`
}
