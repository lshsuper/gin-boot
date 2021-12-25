package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)



//基础控制器
type BaseController struct {
	//基础上下文
	Ctx *gin.Context
}

func (b *BaseController)setContext(c *gin.Context){
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
		methodName=="controllername"||
		methodName=="ok"||
		methodName=="fail"||
		methodName=="result"||
		methodName=="gettraceid"||
		methodName=="gettraceidkey"{
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
	ctrlName:=t[len(t)-1]
	i:=strings.Index(ctrlName,"Controller")
	if i>0{
		ctrlName=ctrlName[:i]
	}
	return ctrlName
}

func (b *BaseController)Ok(data interface{})  {

	b.Ctx.JSON(http.StatusOK,gin.H{
		"data":data,
		"msg":"",
        "code":200,
	})


}

func (b *BaseController)Fail(msg string)  {

	b.Ctx.JSON(http.StatusOK,gin.H{
		"data":nil,
		"msg":msg,
		"code":500,
	})

}

func (b *BaseController)Result(data interface{},code int,msg string)  {

	b.Ctx.JSON(http.StatusOK,gin.H{
		"data":nil,
		"msg":msg,
		"code":500,
	})

}

func (b *BaseController)GetTraceID()string  {

	return b.Ctx.GetHeader(getTraceIDKey())

}

func (b *BaseController)GetTraceIDKey()string  {
    return getTraceIDKey()
}





