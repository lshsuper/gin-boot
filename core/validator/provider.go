package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type provider struct {
	ctx *gin.Context
}

//Register 程序运行时，必须注册
func Register() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for _, tagFunc := range GetValidTags() {
			v.RegisterValidation(tagFunc.String(), tagFunc.ValidFunc())
		}
	}
}

//Form 构造
func Form(ctx *gin.Context) *provider {
	return &provider{ctx: ctx}
}

//CheckJSON 从json标签中验证参数
func (provider *provider) CheckJSON(req IRequest) error {

	if err := req.BeforeCheck(); err != nil {
		return err
	}

	if err := provider.ctx.ShouldBindWith(req, binding.JSON); err != nil {
		return provider.errFormat(req, err, "json")
	}

	if err := req.AfterCheck(); err != nil {
		return err
	}

	return nil

}

//CheckForm 从body-form中验证参数
func (provider *provider) CheckForm(req IRequest) error {

	if err := req.BeforeCheck(); err != nil {
		return err
	}

	if err := provider.ctx.ShouldBindWith(req, binding.Form); err != nil {
		return provider.errFormat(req, err, "form")
	}

	if err := req.AfterCheck(); err != nil {
		return err
	}

	return nil
}

func (provider *provider) errFormat(req IRequest, err interface{}, tag string) error {

	if err == nil {
		return nil
	}

	errArr := err.(validator.ValidationErrors)
	fErr := errArr[0]
	fName := fErr.Field()
	ft := reflect.TypeOf(req).Elem()
	field, _ := ft.FieldByName(fName)
	fTag := field.Tag.Get(tag)

	return req.Check(fErr.Tag(), fTag, fErr.Param())

}
