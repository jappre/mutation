package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func main() {
	db, err := sql.Open("mysql", "root:tiger@tcp(localhost:3306)/user_center")
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println("there is an error1!")
	}
	defer db.Close()

	// Prepare statement for reading data
	stmtOut, err2 := db.Prepare("SELECT * from user")
	if err2 != nil {
		// panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println("there is an error2!")
	}
	var id int
	var email string
	err = stmtOut.QueryRow(1).Scan(&id)
	fmt.Printf("the id is %d, email is %s \n", id, email)
	// for stmtOut.Next() {
	// 	var id int
	// 	var email string
	// 	err = stmtOut.Scan(&id, &email)
	// 	fmt.Printf("the id is %d, email is %s", id, email)
	// }
	defer stmtOut.Close()

	// type squareNum struct {
	// 	id       string
	// 	email    string
	// 	password string
	// }

	// var id int
	rows, err3 := db.Query("SELECT * from user where create_time >'2014-08-28 16:00:00'")
	// err = stmtOut.QueryRow(1).Scan(&id) // WHERE number = 13
	if err3 != nil {
		// panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Printf("there is an error!,%v", err.Error())
	}
	// columns, err := rows.Columns()
	//
	// values := make([]sql.RawBytes, len(columns))
	//
	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }
	// // Fetch rows
	// for rows.Next() {
	// 	// get RawBytes from data
	// 	fmt.Println(values)
	// 	err = rows.Scan(scanArgs...)
	// 	fmt.Println(values)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	// 	}
	//
	// 	// Now do something with the data.
	// 	// Here we just print each column as a string.
	// 	var value string
	// 	for i, col := range values {
	// 		// Here we can check if the value is nil (NULL value)
	// 		if col == nil {
	// 			value = "NULL"
	// 		} else {
	// 			value = string(col)
	// 		}
	// 		fmt.Println(columns[i], ": ", value)
	// 	}
	// 	fmt.Println("-----------------------------------")
	// }
	// if err = rows.Err(); err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	for rows.Next() {
		var valuesofcolumn interface{}
		err = rows.Scan(valuesofcolumn)
		fmt.Printf("the id is %v", valuesofcolumn)
	}
	// fmt.Printf("The square number of 13 is: %d", id)
}
