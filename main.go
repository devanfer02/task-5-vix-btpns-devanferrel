package main

import (
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/router"
)

func main() {
	server := routes.InitRouter()

	server.Run(":8060")
}