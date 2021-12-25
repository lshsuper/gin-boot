package core


type ApiResult struct {
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func Ok(data interface{}) ApiResult {
	return ApiResult{
		Data: data,
		Code:200,
		Msg:"",
	}
}

func Fail(msg string) ApiResult {
	return ApiResult{
		Data: nil,
		Code:500,
		Msg:msg,
	}
}
func Result(code int,data interface{},msg string) ApiResult {
	return ApiResult{
		Data: nil,
		Code:500,
		Msg:msg,
	}
}

