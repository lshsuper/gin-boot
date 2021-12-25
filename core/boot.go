package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type GinBoot struct {
	*gin.Engine
	routeStrict  bool
	addr string
	//读超时设定
	readTimeout  int
	//写超时设定
	writeTimeout int

}

func NewGinBoot(engine *gin.Engine, routeStrict bool, addr string, readTimeout int, writeTimeout int) *GinBoot {
	return &GinBoot{Engine: engine, routeStrict: routeStrict, addr: addr, readTimeout: readTimeout, writeTimeout: writeTimeout}
}




// ServeHTTP (自定义拦截策略)
func (boot *GinBoot) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if !boot.routeStrict{
		req.URL.Path=strings.ToLower(req.URL.Path)
	}

	//gin自带的拦截器
	boot.Engine.ServeHTTP(w,req)

}

func (boot *GinBoot)UseTraceID()*GinBoot {

	boot.Engine.Use(TraceID(getTraceIDKey()))
	return boot
}

func (boot *GinBoot)Use(handler ...gin.HandlerFunc)*GinBoot  {
	boot.Engine.Use(handler...)
	return boot
}


//Register 注册路由
func (boot *GinBoot)Register(fn func() IController)*GinBoot {

	c:=fn()
	t:=reflect.TypeOf(c)
	ctrlName:=c.ControllerName(c)

	for i:=0;i<t.NumMethod();i++{

		methodName:=t.Method(i).Name

		if c.IgnoreMethod(methodName){
			continue
		}

		//判断一下路由是否严格模式
		if !boot.routeStrict{
			methodName=strings.ToLower(methodName)
			ctrlName=strings.ToLower(ctrlName)
		}

		actionUrl:=fmt.Sprintf("%s/%s",ctrlName,methodName)
		//方法设定
		methodType:=c.GetMethodType(methodName)
		boot.Handle(methodType.String(),fmt.Sprintf("/%s",actionUrl), func(context *gin.Context) {
			arr:=strings.Split(context.Request.URL.Path,"/")
			ctrl:=fn()
			ctrl.setContext(context)
			ctrl.CallMethod(ctrl,arr[len(arr)-1])

		})


	}

	return boot

}

//AutoRegister 自动注册所有路由
func (boot *GinBoot)AutoRegister()  {
	//TODO 待实现
}

//Run 启动
func (boot *GinBoot)Run(){

	server:=&http.Server{
		Handler: boot,
		Addr: boot.addr,
	}
	server.ListenAndServe()
}




