package validator

type IRequest interface {
	BeforeCheck() error
	Check(field, tag, param string) error
	AfterCheck() error
}

type BaseResult struct {
}

func (b *BaseResult) BeforeCheck() error {

	return nil

}

func (b *BaseResult) Check(validTag, tag, param string) error {

	return ValidTag(validTag).Message(tag, param)
}

func (b *BaseResult) AfterCheck() error {

	return nil

}
