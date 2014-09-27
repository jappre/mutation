package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func main() {
	db, err := sql.Open("mysql", "root:tiger@tcp(127.0.0.1:3306)/user_center")
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println("there is an error1!")
	}
	defer db.Close()

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT * from user")
	if err != nil {
		// panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println("there is an error2!")
	}
	defer stmtOut.Close()

	// type squareNum struct {
	// 	id       string
	// 	email    string
	// 	password string
	// }

	// var id int
	rows, err := db.Query("SELECT * from user where create_time >'2014-09-01 16:00:00'")
	// err = stmtOut.QueryRow(1).Scan(&id) // WHERE number = 13
	if err != nil {
		// panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Printf("there is an error!,%v", err.Error())
	}
	for rows.Next() {
		var id int
		var email string
		err = rows.Scan(&id, &email)
		fmt.Printf("the id is %d, email is %s", id, email)
	}
	// fmt.Printf("The square number of 13 is: %d", id)
}
