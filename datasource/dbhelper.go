package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hu-bu-la/hubula/conf"
	"github.com/hu-bu-la/hubula/models"
	"github.com/xormplus/xorm"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

//dbLock 互斥锁
var dbLock sync.Mutex

//保持一个
var masterInstance *xorm.Engine
//var slaveInstance *xorm.Engine

//InstanceDbMaster 单例的模式 得到唯一的主库实例
func InstanceDbMaster() *xorm.Engine {
	//如果已经连接 就返回
	if masterInstance != nil {
		return masterInstance
	}
	//如果存在并发
	//加锁 创建之前锁定
	dbLock.Lock()
	//解锁
	defer dbLock.Unlock()
	//也许存在排队的可能
	if masterInstance != nil {
		return masterInstance
	}
	//创建连接
	return NewDbMaster()
}

func dsn() string {
	db_user := conf.Configs().DB["dbMaster"].User
	db_pass := conf.Configs().DB["dbMaster"].Pwd
	db_host := conf.Configs().DB["dbMaster"].Host
	db_port := conf.Configs().DB["dbMaster"].Ports
	db_name := conf.Configs().DB["dbMaster"].Database
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&loc=Asia%2FShanghai"
	return dsn
}

//NewDbMaster 实例化xorm数据库的操作引擎 这个方法是每一次都会实例化一个
func NewDbMaster() *xorm.Engine {
	//连接数据库
	instance, err := xorm.NewEngine(conf.Configs().DB["dbMaster"].DriverName, dsn())
	if err != nil {
		log.Fatal("dbhelper.InstanceDbMaster NewEngine error ", err)
		return nil
	}
	//打印连接信息
	if err = instance.Ping(); err != nil {
		fmt.Println(err)
	}

	//xorm reverse mysql root:root@tcp(127.0.0.1:3306)/hubula?charset=utf8mb4 templates/goxorm
	//同步创建数据表
	err = instance.Sync2(
		new(models.ADemo))

	if err != nil {
		panic(err.Error())
	}

	//调试用的。展示每一条调试语句调试时间
	instance.ShowSQL(true)
	//instance.ShowSQL(false)

	locat, _ := time.LoadLocation("Asia/Shanghai")
	instance.TZLocation = locat

	//返回实例
	masterInstance = instance
	return masterInstance
}


//事务操作
type QuerySession struct {
	Session *xorm.Session
}

var Query *QuerySession

//将map[string]interface{} 分解成表字段和值 进行拼接查询
func Filter(where map[string]interface{}) *xorm.Session {
	db := masterInstance
	Query = new(QuerySession)
	if len(where) > 0 {
		i := 1
		for k, v := range where {
			//fmt.Println(k, v, reflect.TypeOf(v))
			//fmt.Println("?号个数为", strings.Count(k, "?"))
			QuestionMarkCount := strings.Count(k, "?")
			isEmpty := false
			isMap := false
			arrCount := 0
			str := ""
			var arr []string
			switch v.(type) {
			case string:
				//是字符时做的事情
				isEmpty = v == ""
			case int:

			//是整数时做的事情
			case []string :
				isMap = true
				arr = v.([]string)
				arrCount = len(arr)
				isEmpty = arrCount == 0
				for j, val := range arr {
					if j > 0 {
						str += ","
					}
					str += val
				}
			case []int :
				isMap = true
				arrInt := v.([]int)
				arrCount = len(arrInt)
				isEmpty = arrCount == 0
				for j, val := range arrInt {
					if j > 0 {
						str += ","
					}
					str += strconv.Itoa(val)
				}
			}
			if QuestionMarkCount == 0 && isEmpty {
				FilterWhereAnd(db, i, k, "")
			} else if QuestionMarkCount == 0 && !isEmpty {
				//是数组
				if (isMap) {

					FilterWhereAnd(db, i, k, str)
				} else {
					//不是数组
					FilterWhereAnd(db, i, k + " = ?", v)
				}
			} else if QuestionMarkCount == 1 && isEmpty {
				//值为空字符串,不是数组
				FilterWhereAnd(db, i, k, "''")
			} else if QuestionMarkCount == 1 && !isEmpty {
				//是数组
				if isMap {
					//fmt.Println("ArrToStr_key", k)
					//fmt.Println("ArrToStr", str)
					if arrCount > 1 {
						new_q := ""
						for z := 1; z <= arrCount; z++ {
							if z > 1 {
								new_q += ","
							}
							new_q += "?"
						}
						str2 := strings.Replace(k, "?", new_q, -1)
						//fmt.Println("ArrToStr", str)
						//fmt.Println("arr", arr)
						//var inter =arr
						inter := make([]interface{}, arrCount)
						for y, x := range arr {
							inter[y] = x
						}
						FilterWhereAnd(db, i, str2, inter...)
					} else {
						//fmt.Println("22222", str)
						FilterWhereAnd(db, i, k, str)
					}

				} else {
					//不是数组
					//不是数组，有值
					FilterWhereAnd(db, i, k, v)
				}
			} else if QuestionMarkCount > 1 && isEmpty {
				//不是数组，空值
				FilterWhereAnd(db, i, k, "")
			} else if QuestionMarkCount > 1 && !isEmpty && isMap {
				//问号 与  数组相同时
				if QuestionMarkCount == arrCount {
					//不是数组
					FilterWhereAnd(db, i, k, v)
				} else {
					//问号 与  数组不同时
					FilterWhereAnd(db, i, k, str)
				}
			} else {
				fmt.Println("其他还没有收录")
			}
			i++
		}
	} else {
		//初始化
		Query.Session = db.Limit(20, 0)
	}

	return Query.Session
}

//判断是否启用事务操作
func FilterWhereAnd(db *xorm.Engine, i int, key string, value ...interface{}) {
	//fmt.Println("key", key)
	//fmt.Println("value", value)
	//fmt.Println("TypeOf", reflect.TypeOf(value))
	if i == 1 {
		Query.Session = db.Where(key, value...)
	} else {
		Query.Session = Query.Session.And(key, value...)
	}
}