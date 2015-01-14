package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Email    string
	Password string
	Type     int
	Status   int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}
