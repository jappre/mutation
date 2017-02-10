package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=10.0.1.101 user=postgres password=postgres dbname=tlbs sslmode=disable")
	checkErr(err)

	tx, err := db.Begin()
	checkErr(err)

	stmt1, err := tx.Prepare("update tlbs.staff.bs_staff_sign_code set code=$1 where org_id=$2")
	stmt2, err := tx.Prepare("update tlbs.staff.bs_staff_sign_code set code=$1 where org_id=$2")
	checkErr(err)

	res1, err := stmt1.Exec("iuowje", "123")
	res2, err := stmt2.Exec("iuowje", "123224")
	checkErr(err)

	tx.Commit()

	affect1, err := res1.RowsAffected()
	affect2, err := res2.RowsAffected()
	checkErr(err)

	fmt.Println(affect1)
	fmt.Println(affect2)
}

//changeUserId 改UID 8张表
func changeUserID(old_uid string, new_uid string) {
	db, err := sql.Open("postgres", "host=10.0.1.101 user=postgres password=postgres dbname=tlbs sslmode=disable")
	checkErr(err)
	tx, err := db.Begin()
	checkErr(err)

	stmt1, err := tx.Prepare("update tlbs.staff.fy_courier_sign set user_id=$1 where user_id=$2")
	checkErr(err)
	stmt2, err := tx.Prepare("update tlbs.staff.fy_user_video set user_id=$1 where user_id=$2")
	checkErr(err)
	stmt3, err := tx.Prepare("update tlbs.staff.fy_user_record set user_id=$1 where user_id=$2")
	checkErr(err)
	stmt4, err := tx.Prepare("update tlbs.staff.fy_staff_org_struct_staff set user_id=$1 where user_id=$2")
	checkErr(err)
	stmt5, err := tx.Prepare("update tlbs.staff.bs_user_weixin_token set user_id=$1 where user_id=$2")
	checkErr(err)
	stmt6, err := tx.Prepare("update tlbs.staff.bs_user_token set user_id=$1 where user_id=$2")
	checkErr(err)
	stmt7, err := tx.Prepare("update tlbs.staff.bs_user_roles_relation set user_id=$1 where user_id=$2")
	checkErr(err)
	stmt8, err := tx.Prepare("update tlbs.staff.bs_user_account_info set user_id=$1 where user_id=$2")
	checkErr(err)
	res1, err := stmt1.Exec(new_uid, old_uid)
	checkErr(err)
	res2, err := stmt2.Exec(new_uid, old_uid)
	checkErr(err)
	res3, err := stmt3.Exec(new_uid, old_uid)
	checkErr(err)
	res4, err := stmt4.Exec(new_uid, old_uid)
	checkErr(err)
	res5, err := stmt5.Exec(new_uid, old_uid)
	checkErr(err)
	res6, err := stmt6.Exec(new_uid, old_uid)
	checkErr(err)
	res7, err := stmt7.Exec(new_uid, old_uid)
	checkErr(err)
	res8, err := stmt8.Exec(new_uid, old_uid)
	checkErr(err)
	tx.Commit()

	affect1, _ := res1.RowsAffected()
	affect2, _ := res2.RowsAffected()
	affect3, _ := res3.RowsAffected()
	affect4, _ := res4.RowsAffected()
	affect5, _ := res5.RowsAffected()
	affect6, _ := res6.RowsAffected()
	affect7, _ := res7.RowsAffected()
	affect8, _ := res8.RowsAffected()

	fmt.Printf("fy_courier_sign changed row = %d", affect1)
	fmt.Printf("fy_courier_sign changed row = %d", affect2)
	fmt.Printf("fy_courier_sign changed row = %d", affect3)
	fmt.Printf("fy_courier_sign changed row = %d", affect4)
	fmt.Printf("fy_courier_sign changed row = %d", affect5)
	fmt.Printf("fy_courier_sign changed row = %d", affect6)
	fmt.Printf("fy_courier_sign changed row = %d", affect7)
	fmt.Printf("fy_courier_sign changed row = %d", affect8)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
