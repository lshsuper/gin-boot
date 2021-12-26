# gin-boot
#### *说明:对gin框架做的二次封装

#### *简单使用

>1.获取
```
go get github.com/lshsuper/gin-boot

```

>2.定义控制器

```
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

//可以将请求参数全部映射到结构体
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
```

>3.启动

```
boot:= server.New(server.GinBootConf{
		RouteStrict: false,  ////路由严格匹配（true->表示启动路由大小写严格匹配模式|false->表示忽略路由大小写匹配） 
		Addr: ":10086",
	})
	boot.Register(func() core.IController {
		return &Ex01Controller{}
	})
	boot.Run()
	
```

