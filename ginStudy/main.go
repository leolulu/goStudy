package main

import (
	"ginStudy/webserver"
	"log"
)

func main() {
	router := webserver.GetRouter()
	err := router.Run("0.0.0.0:8081")
	if err != nil {
		log.Panic(err)
	}
}
