package handler

import (
	"panorama/server/model"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type RouterHandler struct {
	Hh *gin.Engine
	db model.DBHandler
}

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func MakeHandler() *RouterHandler {
	r := gin.Default()
	rh := &RouterHandler{
		Hh: r,
		db: model.NewDBHandler(),
	}

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		user.POST("signin", rh.signinHandler)
		user.POST("signup", rh.signupHandler)
		post := v1.Group("/post")
		{
			post.POST("img", rh.upLoadImgHandler)
			post.StaticFS("", gin.Dir("", true))
			post.DELETE("img", rh.deleteImgHandler)

			post.GET(":id", rh.getPostHandler)
			post.PATCH("", rh.modifyPostHandler)
			post.POST("", rh.upLoadPostHandler) //contents 동시에 가져와야함
		}
	}
	return rh
}
