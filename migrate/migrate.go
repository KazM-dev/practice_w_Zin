package main

import (
	"fmt"
	"zin-go-rest-api/db"
	"zin-go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
