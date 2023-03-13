package main

import (
	"fmt"
	"os"
	"wikipediaTest/src/repository"
	"wikipediaTest/src/service"
)

func main() {
	fmt.Println("Hello world")
	//time.Sleep(1 * time.Minute)

	if os.Getenv("CREATE_TABLE") == "yes" {
		if os.Getenv("DB_SWITCH") == "on" {
			if err := repository.CreateSchemaAndTable(); err != nil {
				panic(err)
			}
		}
	}
	//time.Sleep(1 * time.Minute)
	service.TelegramBot()
}
