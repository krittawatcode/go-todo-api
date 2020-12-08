package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/krittawatcode/go-todo-api/config"
	"github.com/krittawatcode/go-todo-api/models"
	"github.com/krittawatcode/go-todo-api/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer config.DB.Close()
	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	//setup routes
	r := routes.SetupRouter()
	// running
	r.Run(":8080")
}
