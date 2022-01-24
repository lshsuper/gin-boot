package examples

import (
	"github.com/lshsuper/gin-boot/core"
	"github.com/lshsuper/gin-boot/server"
)

type Ex03Controller struct {
	core.RestController
}

func (c *Ex03Controller) Add() {

	c.Ok("ok add")

}
func (c *Ex03Controller) Get() {

	c.Fail("err get")

}
func (c *Ex03Controller) GetAll() {

	c.Ok("ok getall")

}

func Ex03() {

	boot := server.New(server.GinBootConf{
		RouteStrict: false, //路由严格匹配（忽略大小写的匹配模式）
		Addr:        ":10086",
	})
	ex03 := &Ex03Controller{}
	boot.Register(ex03)
	boot.Run()
}
