package fronted

import (
	"fmt"
	"github.com/hu-bu-la/hubula/conf"
	"github.com/hu-bu-la/hubula/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

var (
	www  =  "www/"
	pathWww = conf.MubanIndexPath + conf.PcPath + www
)
//IndexController
type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

//Get http://localhost:8080/
//返回html页面
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go+iris <a href='http://localhost:8080/cha'>网站</a>"
}

//GetInfo http://localhost:8080/info
//返回json数据
func (c *IndexController) GetInfo() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 200
	rs["msg"] = "成功"
	return rs
}

//GetInfo http://localhost:8080/pass
//返回错误页面
func (c *IndexController) GetPass() {
	c.Ctx.StatusCode(404)
	return
}

//GetInfo http://localhost:8080/test
//返回html模板
func (c *IndexController) GetTest() mvc.Result{
	query := make(map[string]interface{})
	//查询
	data, err := c.ServiceADemo.GetAll(query, []string{}, "time_add,aid desc", 1, 1)
	//错误输出
	if err != nil {
		fmt.Println("err", err)
		//c.Error(err.Error())
		//return
	}
	return mvc.View{
		Name: pathWww + "all.html",
		Data: iris.Map{
			"Title":    "插入和查询",
			"pagings":       data,
		},
		Layout: pathWww + "layout.html",
	}
}