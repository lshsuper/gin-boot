package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/gin-boot/core"
)

type GinBootConf struct {
	RouteStrict  bool
	Addr string
	ReadTimeout int
	WriteTimeout int
}
//New 构造函数
func New(conf GinBootConf) *core.GinBoot {
	r:=gin.New()
	boot:=core.NewGinBoot(r,conf.RouteStrict,conf.Addr,conf.ReadTimeout,conf.WriteTimeout)
	return boot
}


