package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// db, err := sql.Open("postgres", "host=10.0.1.101 user=postgres password=postgres dbname=tlbs sslmode=disable")
	// checkErr(err)
	//
	// tx, err := db.Begin()
	// checkErr(err)
	//
	// stmt1, err := tx.Prepare("update tlbs.staff.bs_staff_sign_code set code=$1 where org_id=$2")
	// stmt2, err := tx.Prepare("update tlbs.staff.bs_staff_sign_code set code=$1 where org_id=$2")
	// checkErr(err)
	//
	// res1, err := stmt1.Exec("iuowje", "123")
	// res2, err := stmt2.Exec("iuowje", "123224")
	// checkErr(err)
	//
	// tx.Commit()
	//
	// affect1, err := res1.RowsAffected()
	// affect2, err := res2.RowsAffected()
	// checkErr(err)
	//
	// fmt.Println(affect1)
	// fmt.Println(affect2)
	GetUserID()
}

//GetUserID 改UID 8张表
func GetUserID() {
	db, err := sql.Open("postgres", "host=182.92.115.196 user=postgres password=feng123 dbname=tlbs sslmode=disable")
	checkErr(err)

	//查询数据
	rows, err := db.Query("select user_id,role from tlbs.staff.fy_framework_staff where status=1 and user_id notnull")
	checkErr(err)

	for rows.Next() {
		var user_id string
		var role string
		err = rows.Scan(&user_id, &role)
		checkErr(err)
		fmt.Printf("user_id = %s,role=%s \n", user_id, role)
		UpdateRoleID(user_id, role)
	}
}

//UpdateRoleID 更新ROLE
func UpdateRoleID(uid string, roleID string) {
	db, err := sql.Open("postgres", "host=182.92.115.196 user=postgres password=feng123 dbname=tlbs sslmode=disable")
	checkErr(err)
	// tx, err := db.Begin()
	// checkErr(err)

	stmt1, err := db.Prepare("update tlbs.staff.bs_user_account_info set job_id=$2 where id=$1")
	checkErr(err)

	res1, err := stmt1.Exec(uid, roleID)
	checkErr(err)

	// tx.Commit()

	affect1, _ := res1.RowsAffected()

	fmt.Printf("bs_user_account_info changed row = %d \n", affect1)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
