package main

import (
	"activeReconBot/service"
	serviceConfig "activeReconBot/service/config"
	"flag"
	"fmt"
	"log"
)

const VERSION = "0.1a"

func main() {
	fmt.Println(fmt.Sprintf("Active Recon Framework API Server v%s - Author: Bastian Muhlhauser @xpl0ited11", VERSION))

	port := flag.Int("p", 3100, "Set the listening port")
	ip := flag.String("b", "0.0.0.0", "Set the listening ip address")

	flag.Parse()

	config := serviceConfig.GetConfig()

	app := &service.App{}

	log.Println("[*] Initiating...")
	app.Initialize(config)

	log.Printf("[*] Binding to %s:%d\n", *ip, *port)
	log.Println("[+] Server started")
	app.Run(fmt.Sprintf("%s:%d", *ip, *port))
}
