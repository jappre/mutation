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

type Real_info struct {
	Id        int
	Real_name string
	User      *User
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Real_info))

	orm.RegisterDataBase("default", "mysql", "root:tiger@tcp(localhost:3306)/user_center?charset=utf8", 30)
}

func main() {
	o := orm.NewOrm()
	// var real_infos []*Real_info
	// qs := o.QueryTable("real_info")
	// num, err := qs.Filter("User__Id", 7749471).All(&real_infos)
	// fmt.Printf("num: %d\n, err: %v \n", 0, qs)

	err = o.Read(&u)
	fmt.Printf("ERR: %v\n, u= %v", err, u)
}
