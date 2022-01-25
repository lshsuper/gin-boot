package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidTag string

const (
	Phone ValidTag = "phone"
	Must  ValidTag = "must"
	Mail  ValidTag = "mail"
	GT    ValidTag = "gt"
	LT    ValidTag = "lt"
	EQ    ValidTag = "eq"
	GTE   ValidTag = "gte"
	LTE   ValidTag = "lte"
	IN    ValidTag="in"
	NEQ   ValidTag="neq"
	LEN  ValidTag="len"
	MAX  ValidTag="max"
	MIN ValidTag="min"
)

func GetValidTags() []ValidTag {
	return []ValidTag{
		Phone, Must, Mail, GT, LT, EQ, GTE, LTE,IN,MAX,MIN,LEN,
	}
}

func (e ValidTag) String() string {
	return string(e)
}

func (e ValidTag) Message(field, param string) error {

	switch e {

	case Phone:
		return errors.New(fmt.Sprintf("[%s]不是正确的手机号", field))
	case Must:
		return errors.New(fmt.Sprintf("[%s]不能为空", field))
	case GT:
		return errors.New(fmt.Sprintf("[%s]必须大于%s", field, param))
	case Mail:
		return errors.New(fmt.Sprintf("[%s]不是正确的邮箱", field))
	case LT:
		return errors.New(fmt.Sprintf("[%s]必须小于%s", field, param))
	case EQ:
		return errors.New(fmt.Sprintf("[%s]必须等于%s", field, param))
	case NEQ:
		return errors.New(fmt.Sprintf("[%s]必须不等于%s", field, param))
	case GTE:
		return errors.New(fmt.Sprintf("[%s]必须大于等于%s", field, param))
	case LTE:
		return errors.New(fmt.Sprintf("[%s]必须小于等于%s", field, param))
	case IN:
		return errors.New(fmt.Sprintf("[%s]必须在%s之内", field, param))
	case LEN:
		return errors.New(fmt.Sprintf("[%s]长度必须等于%s", field, param))
	case MAX:
		return errors.New(fmt.Sprintf("[%s]最大长度等于%s", field, param))
	case MIN:
		return errors.New(fmt.Sprintf("[%s]最小长度等于%s", field, param))
	default:
		return nil
	}
}

//ValidFunc 验证函数匹配
func (e ValidTag) ValidFunc() func(fl validator.FieldLevel) bool {

	switch e {

	case Phone:
		return phoneStrategy
	case Must:
		return mustStrategy
	case GT:
		return gtStrategy
	case LT:
		return ltStrategy
	case EQ:
		return eqStrategy
	case GTE:
		return gteStrategy
	case LTE:
		return lteStrategy
	case Mail:
		return mailStrategy
	case IN:
		return inStrategy
	case NEQ:
		return neqStrategy
	case LEN:
		return lenStrategy
	case MAX:
		return maxStrategy
	case MIN:
		return minStrategy
	default:
		return nil

	}

}
