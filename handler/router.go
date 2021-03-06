package handler

import (
	"io"
	"os"
	"panorama/server/info"
	"panorama/server/model"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type RouterHandler struct {
	Hh *gin.Engine
	db model.DBHandler
}

var client = redis.NewClient(&redis.Options{
	Addr:     info.RedisHost + ":" + info.RedisPort,
	Password: "",
	DB:       0,
})

func MakeHandler() *RouterHandler {
	r := gin.Default()

	gin.DisableConsoleColor()

	rh := &RouterHandler{
		Hh: r,
		db: model.NewDBHandler(),
	}
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.Use(CORSMiddleware())

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		user.POST("signin", rh.signinHandler)
		user.POST("signup", rh.signupHandler)
		user.DELETE("", rh.deleteUserHandler)
		user.GET("signout", rh.signoutHandler)
		post := v1.Group("/project")
		{
			img := post.Group("/img")
			{
				img.POST("upload", rh.upLoadImgHandler)
				img.StaticFS("", gin.Dir("", true))
				img.DELETE("", rh.deleteImgHandler)
			}
			post.GET(":id", rh.getProjectByIdHandler)
			post.GET("", rh.getEntireProjectHandler)
			post.PATCH(":id", rh.modifyProjectHandler)
			post.POST("", rh.upLoadProjectHandler) //contents 동시에 가져와야함
		}
		comment := v1.Group("/comment")
		{
			comment.POST("", rh.uploadCommentHandler)
			comment.GET(":id", rh.getCommentHandler)
		}

	}
	return rh
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Methods", "POST,GET,DELETE")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
