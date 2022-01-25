package core

import "github.com/gin-gonic/gin"

//IController 控制器接口(自定义控制器可以对该接口所有方法做重写)
type IController interface {
	//设定gin的上下文
	setContext(c *gin.Context)
	clone(c IController) IController
	setTraceIDKey(traceIDKey string)

	//CallMethod 动态执行action
	CallMethod(ctrl IController, methodName string)
	//IgnoreMethod 忽略注册的方法策略
	IgnoreMethod(methodName string) bool
	//GetMethodType 获取具体action的请求method策略
	GetMethodType(methodName string) MethodType
	//ControllerName 控制器名称
	ControllerName(ctrl IController) string
	JSON(data interface{})
	GetTraceIDKey() string
}
