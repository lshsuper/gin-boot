package utils

import (
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

// recover错误，转string
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func BuildStruct(t reflect.Type,valMap map[string]interface{})reflect.Value  {

	m:=reflect.New(t).Elem()
	for i:=0;i<m.NumField();i++{
		curField:=t.Field(i)
		fieldName:=curField.Tag.Get("form")
		if v,ok:=valMap[fieldName];ok{
			TrySetValue(m.Field(i),v)
		}
	}
	return m
}
func QSToMap(qs string)(res map[string]interface{}){
        arr:=strings.Split(qs,"&")
        for _,v:=range arr{
        	vArr:=strings.Split(v,"=")
        	res[vArr[0]]=vArr[1]
		}
		return
}

func TrySetValue(fieldVal reflect.Value,val interface{})  {
	switch fieldVal.Kind() {
		case reflect.Bool:
			fieldVal.SetBool(cast.ToBool(val))
	case  reflect.Int:
		fieldVal.SetInt(int64(cast.ToInt(val)))
	case reflect.String:
		fieldVal.SetString(cast.ToString(val))
	}
}
