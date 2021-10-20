package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main() {
	//Open
	//Create Command
	//Pettern Create Table
	//cmd := `CREATE TABLE IF NOT EXISTS persons(
    //            name STRING,
	//			age  INT)`
	//Pettern Create Record
	//cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	////escape by to use "?" and exec Db.Exec(cmd, "taro", 20)

	//Exec the Command

	//row := Db.QueryRow(cmd) get 1 record
	//err := row.Scan(&p.Name, &p.Age) attach row value to argument(model instance)

	//rows := Db.Query(cmd) get all record match cmd

	//なければ作成される
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS persons(
                name STRING,
				age  INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

}
