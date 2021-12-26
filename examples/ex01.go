package examples

import (
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/gin-boot/core"
	"github.com/lshsuper/gin-boot/server"
)

type Ex01Controller struct {
	core.BaseController
}

type GetRequest struct {
	UserID int `json:"user_id" form:"user_id"`
}

func (req GetRequest)Check()error  {
	return nil
}


type AddRequest struct {
	UserID int `json:"user_id" form:"user_id"`
}

func (c *Ex01Controller)Add(req AddRequest)  {


	c.Ok(gin.H{
		"data":req,
	})

}

func (c *Ex01Controller)TestError(req AddRequest)  {

    panic("异常。。。")

	c.Ok(gin.H{
		"data":req,
	})

}

func (c *Ex01Controller)Get(req GetRequest)  {


	c.Ok(gin.H{
		"key":c.GetTraceIDKey(),
		"value":c.GetTraceID(),
		"req":req,
	})

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
		RouteStrict: false,  //路由严格匹配（true->表示启动路由大小写严格匹配模式|false->表示忽略路由大小写匹配）
		Addr: ":10086",
	})
	
	boot.UseTraceID("abc").
		 UseCore().
		 UseRecover(func(msg string,context *gin.Context) interface{} {
		      return map[string]interface{}{
		      	"err":"出异常啦",
			  }
	     })

	boot.Register(func() core.IController {
		return &Ex01Controller{}
	}).Register(func() core.IController {
		return &Ex01ExtController{}
	})
	boot.Run()
}
