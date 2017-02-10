package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type People struct {
	Name string
	ID   string
	Tel  string
}

func main() {
	db, err := sql.Open("postgres", "host=10.0.1.101 user=postgres password=postgres dbname=tlbs sslmode=disable")
	checkErr(err)

	//插入数据
	// stmt, err := db.Prepare("INSERT INTO tlbs.staff.bs_staff_sign_code(org_id,code) VALUES($1,$2) RETURNING org_id")
	// checkErr(err)
	//
	// res, err := stmt.Exec("2123", "212333")
	// checkErr(err)

	//pg不支持这个函数，因为他没有类似MySQL的自增ID
	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)

	//更新数据
	// stmt, err := db.Prepare("update tlbs.staff.bs_staff_sign_code set code=$1 where org_id=$2")
	// checkErr(err)
	//
	// res, err := stmt.Exec("astaxieupdate", "2123")
	// checkErr(err)
	//
	// affect, err := res.RowsAffected()
	// checkErr(err)
	//
	// fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM tlbs.staff.userinfo")
	checkErr(err)

	for rows.Next() {
		var org_id string
		var code string
		var expire time.Time
		err = rows.Scan(&org_id, &code, &expire)
		checkErr(err)
		fmt.Println(org_id)
		fmt.Println(code)

	}
	//
	// //删除数据
	// stmt, err = db.Prepare("delete from tlbs.staff.userinfo where uid=$1")
	// checkErr(err)
	//
	// res, err = stmt.Exec(1)
	// checkErr(err)
	//
	// affect, err = res.RowsAffected()
	// checkErr(err)
	//
	// fmt.Println(affect)
	//
	// db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
