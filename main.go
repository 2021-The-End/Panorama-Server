package main

// @title Panorama Server APi
// @version 1.0.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.email whktjd0109@gmail.com

// @host localhost:3000
// @BasePath /v2
import (
	"os"
	"panorama/server/handler"
)

func main() {
	port := os.Getenv("PORT")
	h := handler.MakeHandler()

	h.Run(":" + port)
}
