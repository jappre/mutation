package main

import (
	"crypto/md5"
  "encoding/hex"
	"database/sql"
	"fmt"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := martini.Classic()
	m.Put("/update/:userid/pwd/:pwd", UpdatePasswd)
	m.Run()
}

//UpdatePasswd 更新密码
/**
 * @api {Put} /update/:userid/pwd/:pwd 更新订单的传递者
 * @apiName UpdatePasswd
 * @apiGroup User
 */
func UpdatePasswd(params martini.Params) string {
	s1 := md5.Sum([]byte(params["pwd"]))
	fmt.Println(hex.EncodeToString(s1[:]))
	s2 := md5.Sum([]byte(hex.EncodeToString(s1[:]) + params["userid"]))
	fmt.Println(hex.EncodeToString(s2[:]))
	s3 := Substr(hex.EncodeToString(s2[:]), len(hex.EncodeToString(s2[:]))-20, 20)
	fmt.Println(s3)

	db, err := sql.Open("mysql", "fengempire:es#5l%iq7n*b2gs@tcp(123.56.94.138:3306)/empire")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.Exec("UPDATE user SET password=? WHERE staff_num = ?", s3, params["userid"])
	return "修改成功"
}

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
