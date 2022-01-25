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

type BootServer struct {
	*gin.Engine
	routeStrict bool
	addr        string
	//读超时设定
	readTimeout int
	//写超时设定
	writeTimeout int
	traceIDKey   string
}

func NewGinBoot(engine *gin.Engine, routeStrict bool, addr string, readTimeout int, writeTimeout int) *BootServer {
	return &BootServer{Engine: engine, routeStrict: routeStrict, addr: addr, readTimeout: readTimeout, writeTimeout: writeTimeout}
}

// ServeHTTP (自定义拦截策略)
func (boot *BootServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if !boot.routeStrict {
		req.URL.Path = strings.ToLower(req.URL.Path)
	}

	//gin自带的拦截器
	boot.Engine.ServeHTTP(w, req)

}

func (boot *BootServer) UseTraceID(traceIDKey string) *BootServer {

	if len(traceIDKey) <= 0 {
		traceIDKey = defaultTraceIDKey
	}

	boot.traceIDKey = traceIDKey

	boot.Engine.Use(func(context *gin.Context) {

		traceId := context.GetHeader(traceIDKey)
		if len(traceId) <= 0 {
			traceId = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
			context.Request.Header.Add(traceIDKey, traceId)
		}
		context.Next()
	})
	return boot
}

type BootCoreConf struct {
	AllowOrigin   string
	AllowHeaders  string
	AllowMethods  string
	ExposeHeaders string
}

func (boot *BootServer) UseCore(conf *BootCoreConf) *BootServer {

	//不填就是默认配置
	if conf == nil {
		conf = defaultBootCore
	}

	boot.Engine.Use(func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", conf.AllowOrigin)
		context.Header("Access-Control-Allow-Headers", conf.AllowHeaders)
		context.Header("Access-Control-Allow-Methods", conf.AllowMethods)
		context.Header("Access-Control-Expose-Headers", conf.ExposeHeaders)
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

func (boot *BootServer) UseRecover(fn func(msg string, context *gin.Context) interface{}) *BootServer {

	if fn == nil {
		fn = func(msg string, context *gin.Context) interface{} {
			return Fail(msg)
		}
	}

	boot.Engine.Use(func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(http.StatusOK, fn(utils.ErrorToString(r), context))
				//终止当前接口
				context.Abort()
			}

		}()
		//加载完 defer recover，继续后续接口调用
		context.Next()
	})
	return boot
}

func (boot *BootServer) Use(handler ...gin.HandlerFunc) *BootServer {
	boot.Engine.Use(handler...)
	return boot
}

//Register 注册路由
func (boot *BootServer) Register(controllers ...IController) *BootServer {

	for _, controller := range controllers {
		boot.register(boot.Engine, controller)
	}

	return boot

}

func (boot *BootServer) register(e *gin.Engine, controller IController) {
	t := reflect.TypeOf(controller)
	ctrlName := controller.ControllerName(controller)

	for i := 0; i < t.NumMethod(); i++ {

		methodName := t.Method(i).Name

		if controller.IgnoreMethod(methodName) {
			continue
		}

		//判断一下路由是否严格模式
		if !boot.routeStrict {
			methodName = strings.ToLower(methodName)
			ctrlName = strings.ToLower(ctrlName)
		}

		actionUrl := fmt.Sprintf("%s/%s", ctrlName, methodName)
		//方法设定
		methodType := controller.GetMethodType(methodName)

		e.Handle(methodType.String(), fmt.Sprintf("/%s", actionUrl), func(context *gin.Context) {
			arr := strings.Split(context.Request.URL.Path, "/")
			ctrl := controller
			ctrl.setContext(context)
			ctrl.setTraceIDKey(boot.traceIDKey)
			ctrl.CallMethod(ctrl, arr[len(arr)-1])
		})

	}

}

//AutoRegister 自动注册所有路由
func (boot *BootServer) AutoRegister() {
	//TODO 待实现
}

//Run 启动
func (boot *BootServer) Run() {

	server := &http.Server{
		Handler: boot,
		Addr:    boot.addr,
	}
	server.ListenAndServe()
}
