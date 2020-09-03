package main

import (
	"database/sql"
	"fmt"
	"time"
)

type Employee struct {
	id         int
	first_name string
	city       string
}

func initialisingDataBase() {

	db_driver := "mysql"
	db_username := "lambda"
	db_password := "deepspace"
	db_name = "goBlog"

	db_connection, err := sql.Open(db_driver, db_username+":"+db_password+"@/"+db_name)

	if err != nil {

		panic(err.Error())

	}

	//important setting section
	db_connection.SetConnMaxLifetime(time.Minute * 3)
	db_connection.SetMaxOpenConns(10)
	db_connection.SetMaxIdleConns(10)

	return db_connection

}

func main() {
	db_connectionIndex := initialisingDataBase()

	sel_db, err := db_connectionIndex.Query("SELECT * FROM employee ORDER BY id DESC")

	if err != nil {

		panic(err.Error())
	}

	data_emp := Employee{}
	data_res := []Employee{}

	for sel_db.Next() {
		var id int
		var first_name, city string

		err := sel_db.Scan(&id, &first_name, &city)

		if err != nil {

			panic(err.Error())
		}

		data_emp.id = id
		data_emp.first_name = first_name
		data_emp.city = city

		data_res = append(data_emp, data_res)

		defer db_connectionIndex.Close()

	}

	fmt.Println(data_res)
}
