package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lshsuper/gin-boot/core"
)

type BootConf struct {
	RouteStrict  bool
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

//New 构造函数
func New(conf BootConf) *core.BootServer {
	r := gin.New()
	boot := core.NewGinBoot(r, conf.RouteStrict, conf.Addr, conf.ReadTimeout, conf.WriteTimeout)
	return boot
}
