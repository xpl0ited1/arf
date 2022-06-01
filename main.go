package main

import (
	"activeReconBot/service"
	"activeReconBot/service/config"
	"fmt"
	"log"
)

func main() {
	port := 3100

	config := config.GetConfig()

	app := &service.App{}

	log.Println("[*]Initiating...")
	app.Initialize(config)

	log.Printf("[+]Started server at port %d\n", port)
	app.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
