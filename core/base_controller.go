package core

import (
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/gin-boot/utils"
	"net/http"
	"reflect"
	"strings"
)

//BaseController 基础控制器
type BaseController struct {
	//基础上下文
	Ctx        *gin.Context
	traceIDKey string
}

func (b *BaseController) setContext(c *gin.Context) {
	b.Ctx = c
}

//CallMethod 执行具体的方法
func (b *BaseController) CallMethod(ctrl IController, methodName string) {

	curCtrl := reflect.ValueOf(ctrl)
	t := curCtrl.Type()
	for i := 0; i < curCtrl.NumMethod(); i++ {

		curMethodName := t.Method(i).Name
		if strings.ToLower(curMethodName) == strings.ToLower(methodName) {

			curMethod := curCtrl.MethodByName(curMethodName)
			if curMethod.Type().NumIn() <= 0 {
				curMethod.Call(nil)
				break
			}

			ps := make([]reflect.Value, 0)
			pType := curMethod.Type().In(0)
			keyMap := make(map[string]interface{}, 0)
			switch b.Ctx.Request.Method {
			case GET.String():
				b.Ctx.Request.ParseForm()
			case POST:
				fallthrough
			case PUT:
				fallthrough
			case DELETE:
				b.Ctx.Request.ParseMultipartForm(32 << 20)
			}
			for k, v := range b.Ctx.Request.Form {
				keyMap[k] = v[0]
			}
			for k, v := range b.Ctx.Request.Form {
				keyMap[k] = v[0]
			}
			vType := utils.BuildStruct(pType, keyMap)
			ps = append(ps, vType)

			curMethod.Call(ps)
			break
		}

		//b.Ctx.BindWith()

	}

}

//IgnoreMethod 忽略注册方法
func (b *BaseController) IgnoreMethod(methodName string) bool {

	methodName = strings.ToLower(methodName)
	if methodName == "setcontext" ||
		methodName == "callmethod" ||
		methodName == "ignoremethod" ||
		methodName == "getmethodtype" ||
		methodName == "controllername" ||
		methodName == "ok" ||
		methodName == "fail" ||
		methodName == "result" ||
		methodName == "gettraceid" ||
		methodName == "gettraceidkey" ||
		methodName == "json" {
		return true
	}
	return false
}

func (b *BaseController) GetMethodType(methodName string) MethodType {

	//可以设定默认的判定规则,也可以重写该方法自定义规则
	methodName = strings.ToLower(methodName)
	if strings.Index(methodName, "add") >= 0 ||
		strings.Index(methodName, "update") >= 0 ||
		strings.Index(methodName, "edit") >= 0 ||
		strings.Index(methodName, "delete") >= 0 ||
		strings.Index(methodName, "edit") >= 0 {
		return POST
	}
	return GET

}

//ControllerName 控制器名称
func (b *BaseController) ControllerName(ctrl IController) string {

	curCtrl := reflect.ValueOf(ctrl)
	t := strings.Split(curCtrl.Type().String(), ".")
	ctrlName := t[len(t)-1]
	i := strings.Index(ctrlName, "Controller")
	if i > 0 {
		ctrlName = ctrlName[:i]
	}
	return ctrlName
}

func (b *BaseController) Ok(data interface{}) {

	b.Ctx.JSON(http.StatusOK, Ok(data))

}

func (b *BaseController) Fail(msg string) {
	b.Ctx.JSON(http.StatusOK, Fail(msg))
}

func (b *BaseController) Result(data interface{}, code int, msg string) {

	b.Ctx.JSON(http.StatusOK, Result(code, data, msg))

}

func (b *BaseController) JSON(data interface{}) {
	b.Ctx.JSON(http.StatusOK, data)
}

func (b *BaseController) GetTraceID() string {
	return b.Ctx.GetHeader(b.traceIDKey)
}

func (b *BaseController) setTraceIDKey(traceIDKey string) {
	b.traceIDKey = traceIDKey
}
func (b *BaseController) GetTraceIDKey() string {
	return b.traceIDKey
}
