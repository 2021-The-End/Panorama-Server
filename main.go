package main

import (
	"os"
	"panorama/server/handler"
)

func main() {
	port := os.Getenv("PORT")
	rh := handler.MakeHandler()

	rh.Close()

	rh.Run(":" + port)
}
