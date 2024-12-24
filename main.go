package main

import (
	_ "catproject/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.BConfig.Listen.Graceful = false
}

func main() {
	beego.Run()
}
