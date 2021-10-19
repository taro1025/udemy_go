package main

import (
	"log"
	"practiceGo/todo_app/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.Port)

	log.Println("test")
}
