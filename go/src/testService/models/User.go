package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int `orm:"auto"`
	Title string
}

type Comment struct {
	Id      int `orm:"auto"`
	Content string
	Post    *Post `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Post), new(Comment))

	orm.RegisterDataBase("default", "mysql", "root:tiger@tcp(localhost:3306)/bbs?charset=utf8", 30)
}

func main() {
	o := orm.NewOrm()

	// post := Post{Title: "weifang"}
	// id, err := o.Insert(&post)
	// fmt.Printf("ID: %d, ERR: %v \n", id, err)

	var comments []*Comment
	qs := o.QueryTable("comment")
	// num, err := qs.Filter("content__isnull", true).All(&comments)
	num, err := qs.All(&comments)
	fmt.Printf("num = %d, err = %v", num, err)

	// user.Email = "testforgo7@test123gezbox.com"
	// num, err := o.Update(&user)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	// u := User{Id: user.Id}
	// err = o.Read(&u)
	// fmt.Printf("ERR: %v\n, u= %v", err, u)
	//
	// num, err = o.Delete(&u)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
