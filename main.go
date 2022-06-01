package main

import (
	appConfig "activeReconBot/config"
	"activeReconBot/service"
	serviceConfig "activeReconBot/service/config"
	"fmt"
	"log"
)

func main() {
	port := 3100

	config := serviceConfig.GetConfig()

	app := &service.App{}

	log.Println("[*]Initiating...")
	app.Initialize(config)

	log.Printf("[+]Started server at port %d\n", port)
	app.Run(fmt.Sprintf("0.0.0.0:%d", port))
	fmt.Println(appConfig.SHODAN_TOKEN)
}
