package core

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)



//基础控制器
type BaseController struct {
	//基础上下文
	Ctx *gin.Context
}

func (b *BaseController)SetContext(c *gin.Context){
   b.Ctx=c
}


//CallMethod 执行具体的方法
func (b *BaseController)CallMethod(ctrl IController,methodName string) {

	curCtrl:=reflect.ValueOf(ctrl)
	t:=curCtrl.Type()
	for i:=0;i<=curCtrl.NumMethod();i++{

		curMethodName:=t.Method(i).Name
		if strings.ToLower(curMethodName)==strings.ToLower(methodName){
			curCtrl.MethodByName(curMethodName).Call(nil)
			break
		}


	}

}

func (b *BaseController)IgnoreMethod(methodName string)bool  {

	methodName=strings.ToLower(methodName)
	if methodName=="setcontext"||
		methodName=="callmethod"||
		methodName=="ignoremethod"||
		methodName=="getmethodtype"||
		methodName=="controllername"{
		return true
	}
	return false
}

func (b *BaseController)GetMethodType(methodName string) MethodType {

	//可以设定默认的判定规则,也可以重写该方法自定义规则
	methodName=strings.ToLower(methodName)
	if strings.Index(methodName,"add")>=0||
		strings.Index(methodName,"update")>=0||
		strings.Index(methodName,"edit")>=0||
		strings.Index(methodName,"delete")>=0||
		strings.Index(methodName,"edit")>=0{
		return POST
	}
	return GET

}

//ControllerName 控制器名称
func (b *BaseController)ControllerName(ctrl IController)string  {
	curCtrl:=reflect.ValueOf(ctrl)
	t:=strings.Split(curCtrl.Type().String(),".")
	ctrlName:=strings.TrimRight(t[len(t)-1],"Controller")
	return ctrlName
}




