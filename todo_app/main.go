package main

import (
	"log"
	"practiceGo/todo_app/app/models"
	"practiceGo/todo_app/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.Port)

	log.Println(models.Db)
	u := &models.User{}
	u.Name = "tes"
	u.Email = "test"
	u.PassWord = "tlep"
	fmt.Println(u)

	u.CreateUser()
}
