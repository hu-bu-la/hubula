package routes

import (
	"github.com/hu-bu-la/hubula/bootstrap"
	"github.com/hu-bu-la/hubula/services"
	"github.com/hu-bu-la/hubula/web/controllers/fronted"
	"github.com/kataras/iris/v12/mvc"
)

// Configure 和 bootstrap 里定义的一样
func Configure(b *bootstrap.Bootstrapper) {
	//主要是把indexcontrollers放进去 里面定义了很多service
	ademoService := services.NewADemoService()


	//用mvc创建一个新的路径
	index := mvc.New(b.Party("/"))
	//把 Service 都注册进去
	index.Register(ademoService)
	//路径发给Handle
	index.Handle(new(fronted.IndexController))
}