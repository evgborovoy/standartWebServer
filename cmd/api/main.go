package main

import (
	"log"

	"github.com/evgborovoy/StandardWebServer/internal/app/api"
)

var ()

func init() {

}

func main() {
	log.Println("start")
	// server instance initializations
	server := api.New()

	//api server start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
