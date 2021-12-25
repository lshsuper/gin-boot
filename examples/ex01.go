package examples

import (
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/gin-boot/core"
	"github.com/lshsuper/gin-boot/server"
	"net/http"
)

type MyController struct {
	core.BaseController
}

func (c *MyController)Add()  {

	c.Ctx.JSON(http.StatusOK,gin.H{"data":"ok get"})

}
func (c *MyController)Get()  {

	c.Ctx.JSON(http.StatusOK,gin.H{"data":"ok get"})

}
func (c *MyController)GetAll()  {

	c.Ctx.JSON(http.StatusOK,gin.H{"data":"ok GetAll"})

}

func Ex01(){

	boot:= server.New(server.GinBootConf{
		RouteStrict: false,  //路由严格匹配（忽略大小写的匹配模式）
		Addr: ":10086",
	})
	boot.Register(func() core.IController {
		return &MyController{}
	})
	boot.Run()
}
