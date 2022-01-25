package core

const defaultTraceIDKey = "trace_id"

var defaultBootCore = &BootCoreConf{
	AllowMethods:  "*",
	AllowHeaders:  "access-control-allow-origin,content-Type,AccessToken,X-CSRF-Token, Authorization, Token",
	AllowOrigin:   "*",
	ExposeHeaders: "content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, content-Type",
}
