package main

import (
	"app/internal/application"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var storage *sql.DB

func main() {
	// env
	// ...

	// app
	// - config
	app := application.NewApplicationDefault("", storage)
	// - tear down
	defer app.TearDown()
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	dataSource := "root:root@tcp(localhost:3306)/my_db?parseTime=true"

	var err error
	storage, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err := storage.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}
