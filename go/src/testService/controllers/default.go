package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// model "testService/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplNames = "index.tpl"
	c.Ctx.WriteString("hello")
}

func (c *MainController) Post() {

	var ob interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// objectid := models.AddOne(ob)
	// c.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	// c.ServeJson()
	fmt.Println(ob)
	c.Ctx.WriteString("helo, post")
}

// func (c *MainController) ModifyEmail() {
// 	o := orm.NewOrm()
//
// 	// post := Post{Title: "weifang"}
// 	// id, err := o.Insert(&post)
// 	// fmt.Printf("ID: %d, ERR: %v \n", id, err)
//
// 	var comments []*Comment
// 	qs := o.QueryTable("comment")
// 	// num, err := qs.Filter("content__isnull", true).All(&comments)
// 	num, err := qs.All(&comments)
// 	fmt.Printf("num = %d, err = %v", num, err)
//
// 	// user.Email = "testforgo7@test123gezbox.com"
// 	// num, err := o.Update(&user)
// 	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
// 	//
// 	// u := User{Id: user.Id}
// 	// err = o.Read(&u)
// 	// fmt.Printf("ERR: %v\n, u= %v", err, u)
// 	//
// 	// num, err = o.Delete(&u)
// 	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
// }
