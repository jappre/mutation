package main

import (
	_ "testService/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

