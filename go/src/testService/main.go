package main

import (
	"github.com/astaxie/beego"
	_ "testService/routers"
)

func main() {
	beego.Run()
}
