package examples

import (
	"github.com/lshsuper/gin-boot/core"
	"github.com/lshsuper/gin-boot/server"
)

type Ex01Controller struct {
	core.BaseController
}

func (c *Ex01Controller)Add()  {

	c.Ok("ok add")

}
func (c *Ex01Controller)Get()  {

	c.Ok("ok get")

}
func (c *Ex01Controller)GetAll()  {

	c.Ok("ok getall")

}

type Ex01ExtController struct {
	core.BaseController
}

func (c *Ex01ExtController)GetAll()  {

	c.Ok("ok getall")

}

func Ex01(){

	boot:= server.New(server.GinBootConf{
		RouteStrict: false,  //路由严格匹配（忽略大小写的匹配模式）
		Addr: ":10086",
	})
	boot.Register(func() core.IController {
		return &Ex01Controller{}
	}).Register(func() core.IController {
		return &Ex01ExtController{}
	})
	boot.Run()
}
