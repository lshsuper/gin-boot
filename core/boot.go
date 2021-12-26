package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/gin-boot/utils"
	uuid "github.com/satori/go.uuid"
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

func (boot *GinBoot)UseTraceID(traceIDKey string)*GinBoot {

	if len(traceIDKey)<=0{
		traceIDKey=getTraceIDKey()
	}else{
		setTraceIDKey(traceIDKey)
	}

	boot.Engine.Use(func(context *gin.Context) {

		traceId:=context.GetHeader(traceIDKey)
		if len(traceId)<=0{
			traceId=strings.ReplaceAll(uuid.NewV4().String(),"-","")
			context.Request.Header.Add(traceIDKey,traceId)
		}
		context.Next()
	})
	return boot
}

func (boot *GinBoot)UseCore()*GinBoot  {

	boot.Engine.Use(func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "access-control-allow-origin,content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		context.Next()
	})
	return boot
}

func (boot *GinBoot)UseRecover(fn func(msg string,context *gin.Context)interface{})*GinBoot  {

	if fn==nil{
		fn= func(msg string,context *gin.Context) interface{} {
			return Fail(msg)
		}
	}

	boot.Engine.Use(func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(http.StatusOK, fn(utils.ErrorToString(r),context))
				//终止当前接口
				context.Abort()
			}

		}()
		//加载完 defer recover，继续后续接口调用
		context.Next()
	})
	return boot
}

func (boot *GinBoot)Use(handler ...gin.HandlerFunc)*GinBoot  {
	boot.Engine.Use(handler...)
	return boot
}

type MyRequest struct {
	UserID int `json:"user_id" form:"user_id"`
}

//Register 注册路由
func (boot *GinBoot)Register(fns ...func() IController)*GinBoot {

	for _,fn:=range fns{
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




