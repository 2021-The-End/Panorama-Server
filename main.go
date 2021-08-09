package main

import (
	"log"
	"panorama/server/handler"
	"panorama/server/info"
)

func main() {
	port := info.AppPort
	rh := handler.MakeHandler()

	log.Print("Start App")

	rh.Hh.Run(":" + port)
}
