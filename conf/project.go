package conf

import "time"

//SysTimeform 时间
const SysTimeform = "2006-01-02 15:04:05"

//SysTimeformShort 时间
const SysTimeformShort = "2006-01-02"

//IndexPath 前台模板路径
var MubanIndexPath = "fronted/"

//AdminPath 后台模板路径
var MubanAdminPath = "backend/"

//PcPath PC模板路径
var PcPath = "pc/"

//WapPath WAP模板路径
var WapPath = "wap/"

//AutoPath AUTO模板路径
var AutoPath = "auto/"

//RunningCrontabService 是否需要启动全局计划任务服务
var RunningCrontabService = false

//SysTimeLocation 中国时区
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

//SignSecret ObjSalesign 签名密钥
var SignSecret = []byte("0123456789abcdef")

//CookieSecret cookie中的加密验证密钥
var CookieSecret = "hellolottery"

