package examples

import (
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/gin-boot/core"
	"github.com/lshsuper/gin-boot/server"
	"net/http"
)

type MyBaseController struct {
	core.BaseController
}

func (b *MyBaseController)Ok(data interface{})  {

	b.Ctx.JSON(http.StatusOK,gin.H{
		"data":data,
		"msg":"",
		"success":true,
	})


}

func (b *MyBaseController)Fail(msg string)  {

	b.Ctx.JSON(http.StatusOK,gin.H{
		"data":nil,
		"msg":msg,
		"success":false,
	})

}


type Ex02Controller struct {
	MyBaseController
}



func (c *Ex02Controller)Add()  {

	c.Ok("ok add")

}
func (c *Ex02Controller)Get()  {

	c.Ok("ok get")

}
func (c *Ex02Controller)GetAll()  {

	c.Ok("ok getall")

}

func Ex02(){

	boot:= server.New(server.GinBootConf{
		RouteStrict: false,  //路由严格匹配（忽略大小写的匹配模式）
		Addr: ":10086",
	})
	boot.Register(func() core.IController {
		return &Ex02Controller{}
	})
	boot.Run()
}