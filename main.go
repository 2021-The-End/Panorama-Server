package main

import (
	"log"
	"os"
	"panorama/server/handler"
)

func main() {
	port := os.Getenv("PORT")
	rh := handler.MakeHandler()

	log.Print("Start App")

	rh.Hh.Run("127.0.0.1:" + port)
}
