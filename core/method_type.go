package core


type MethodType string

const(
	GET    MethodType ="GET"
	POST              ="POST"
	PUT               ="PUT"
	DELETE            ="DELETE"
)

func (e MethodType)String()string  {
	return string(e)
}

