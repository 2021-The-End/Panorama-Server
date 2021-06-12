package handler

import (
	_ "panorama/server/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func MakeHandler() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		users.GET("", getUsersListHandler)
		users.GET(":id", getUserHandler)
		users.POST("", createUsersHandler)
		users.PUT(":id", updateUsersHandler)

	}
	r.GET("/", indexHandler)
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
