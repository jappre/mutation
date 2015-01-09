package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int
	Email      string `orm:"size(100)"`
	Password   string
	Status     int
	CreateTime string
	Tel        string
}

func init() {
	orm.RegisterModel(new(User))

	orm.RegisterDataBase("default", "mysql", "root:tiger@tcp(localhost:3306)/user_center?charset=utf8", 30)
}

func main() {
	o := orm.NewOrm()

	user := User{Email: "testforgo6@test123gezbox.com", Status: 1, Tel: "18668886667"}
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v \n", id, err)

	user.Email = "testforgo7@test123gezbox.com"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n, u= %v", err, u)

	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
