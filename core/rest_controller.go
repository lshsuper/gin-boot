package core

import (
	"net/http"
	"strings"
)

//RestController
type RestController struct {
	BaseController
}

func (b *RestController) Ok(data interface{}) {

	b.Ctx.JSON(http.StatusOK, data)

}

//GetMethodType 严格rest模式
func (b *RestController) GetMethodType(methodName string) MethodType {

	//可以设定默认的判定规则,也可以重写该方法自定义规则
	methodName = strings.ToLower(methodName)
	if strings.Index(methodName, "add") >= 0 {
		return POST
	}

	if strings.Index(methodName, "update") >= 0 || strings.Index(methodName, "edit") >= 0 {
		return PUT
	}

	if strings.Index(methodName, "delete") >= 0 || strings.Index(methodName, "del") >= 0 {
		return DELETE
	}

	return GET

}

//IgnoreMethod 忽略注册方法
func (b *RestController) IgnoreMethod(methodName string) bool {

	methodName = strings.ToLower(methodName)
	if res := b.BaseController.IgnoreMethod(methodName); res {
		return true
	}

	if methodName == "unauthorized" ||
		methodName == "badrequest" ||
		methodName == "notfound" ||
		methodName == "statuscode" {
		return true
	}

	return false
}

func (b *RestController) Fail(data interface{}) {
	b.Ctx.JSON(http.StatusInternalServerError, data)
}

func (b *RestController) Unauthorized() {
	b.Ctx.JSON(http.StatusUnauthorized, nil)
}

func (b *RestController) BadRequest() {
	b.Ctx.JSON(http.StatusBadRequest, nil)
}

func (b *RestController) NotFound() {
	b.Ctx.JSON(http.StatusBadRequest, nil)
}

func (b *RestController) StatusCode(code int) {
	b.Ctx.JSON(code, nil)
}
