package conf

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	cfg *Config
	once sync.Once
	cfgLock = new(sync.RWMutex)
)

var Rds = `
# 全局信息
title = "TOML格式配置文件示例"
# 网站常量
[webstatic]
    #后台静态文件路径
    assets = "./assets/"
    #前台静态文件路径
    public = "./public/"
    #icon图标名称
    favicon = "favicon.ico"

# 网站变量
[webvar]
    #模板目录
    views = "./views"

# 应用信息
[app]
    name = ""
	owner  = ""
    author = "史布斯"
    release = 2021-01-28T08:18:00Z   # 时间
    port = 8080
    organization = "Mafool"
    mark = "第一行\n第二行."            # 换行

# 数据库配置
# root:root@tcp(127.0.0.1:3306)/lottery?charset=utf-8
[mysql]
    [mysql.dbMaster]
        driverName = "mysql"
        host = "127.0.0.1"
        ports = 3306     # 数组
        user = "root"
        pwd = "root"
        database = "db"
        isRunning = true        # 状态 是否正常运行
        connection_max = 5000
        # enabled = true
    [mysql.dbSlave]
        driverName = "mysql"
        host = "192.168.1.1"
        ports = 3306     # 数组
        user = "root"
        pwd = "123456"
        database = "hubula"
        isRunning = true        # 状态 是否正常运行
        connection_max = 5000
        # enabled = true

# Redis主从                           # 字典对象
[redis]
    [redis.master]
        host = "10.0.0.1"
        port = 6379
    [redis.slave]
        host = "10.0.0.1"
        port = 6380

# 二维数组
[releases]
release = ["dev", "test", "stage", "prod"]
tags = [["dev", "stage", "prod"],[2.2, 2.1]]


# 公司信息                             #对象嵌套
[company]
    name = "xx科技"
[company.detail]
    type = "game"
    addr = "北京朝阳"
    icp = "030173"
[[song]]
name = "天路"
duration = "4m49s"

[[song]]
name = "忘情水"
duration = "8m03s"
`

func Configs() *Config {
	once.Do(ReloadConfig)
	cfgLock.RLock()
	defer cfgLock.RUnlock()
	return cfg
}

func ReloadConfig() {
	//config := new(Config)
	//读取文件
	//filePath, err := filepath.Abs("../conf/gateway.toml")
	//if err != nil {
	//	fmt.Println("load config error: ", err)
	//	panic(err)
	//}
	//fmt.Printf("parse toml file once. filePath: %s\n", filePath)
	//if _ , err := toml.DecodeFile(filePath, &cfg); err != nil {
	//	fmt.Println("Para config failed: ", err)
	//	panic(err)
	//}

	//读取字符串
	if _, err := toml.Decode(Rds, &cfg); err != nil {
		panic(err)
	}

	cfgLock.Lock()
	defer cfgLock.Unlock()
	//cfg = cfg
}

//windows下使用endless报错：undefined: syscall.SIGUSR1
//windows 下的信号没有 SIGUSR1、SIGUSR2 等，为了不轻易的抛弃 windows 环境 (没有勇气换 mac)。目前用了一个 ditty 的办法解决：
//在 go 的安装目录修改 Go\src\syscall\types_windows.go，增加如下代码：
//
//var signals = [...]string{
//	// 这里省略N行。。。。
//
//	/** 兼容windows start */
//	16: "SIGUSR1",
//	17: "SIGUSR2",
//	18: "SIGTSTP",
//	/** 兼容windows end */
//}
//
///** 兼容windows start */
//func Kill(...interface{}) {
//return;
//}
//const (
//SIGUSR1 = Signal(0x10)
//SIGUSR2 = Signal(0x11)
//SIGTSTP = Signal(0x12)
//)
/** 兼容windows end */

func init() {
	//热更新配置可能有多种触发方式，这里使用系统信号量sigusr1实现
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			ReloadConfig()
			log.Println("Reloaded config")
		}
	}()
}


//fmt.Printf("全局信息: %+v\n\n", conf.Configs().Title)
//
//fmt.Printf("App信息：%+v\n\n", conf.Configs().App)
//
//fmt.Printf("Mysql配置：%+v\n\n", conf.Configs().DB)
//
//fmt.Printf("版本信息：%+v\n\n", conf.Configs().Releases)
//
//fmt.Printf("Redis主从：%+v\n\n", conf.Configs().Redis)
//
//fmt.Printf("企业信息：%+v\n\n", conf.Configs().Company)
//
//fmt.Printf("信息：%+v\n\n", conf.Configs().Song)

