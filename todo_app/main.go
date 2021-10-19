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
}
