package main

import (
	"os"
	"fmt"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load();

	if err != nil {
		panic(err)
	}

	server := routes.InitRouter()
	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"));
	
	server.Run(port)
}