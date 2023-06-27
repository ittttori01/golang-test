package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Test struct{
	id int
	code string
}

func mysqlConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:roqkfroqkf@tcp(127.0.0.1:3306)/go")
	if err != nil || db.Ping() != nil { 
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3) 
	db.SetMaxOpenConns(10) 
	db.SetMaxIdleConns(10) 
	return db
}


func findALL() []Test{
	conn := mysqlConnect()
	defer conn.Close()
	rows, err := conn.Query("SELECT * FROM test")
	if err!=nil{
		panic(err)
	}
	defer rows.Close()
	test := []Test{}
	for rows.Next() {
        var t Test
        rows.Scan(&t.id, &t.code)
        test = append(test, t)
    }
	return test
}

func transactionSave(code string) int{
	conn := mysqlConnect()
	defer conn.Close()
	tx, err := conn.Begin()
	if err !=nil{
		tx.Rollback()
		panic(err)
	}

	query := fmt.Sprint("INSERT INTO test (`code`) VALUES ('", code, "')")
	res, err := tx.Exec(query)
	if err!=nil{
        tx.Rollback()
        panic(err)
    }

	lastId,err := res.LastInsertId()
	fmt.Println("========INSERT RESULT=========")
	fmt.Println(res.LastInsertId())
	if err!=nil{
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	return int(lastId)
}


func main() {
	fmt.Println("db crud")
	test := findALL()
	fmt.Println(test)
	lastId := transactionSave("TESTCODE1")
	fmt.Println("SAVED",lastId)
}