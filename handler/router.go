package handler

import (
	"panorama/server/model"
	"panorama/server/utils"

	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	gin.Engine
	db model.DBHandler
	md utils.MiddleWare
}

func MakeHandler() *RouterHandler {
	r := gin.Default()

	rh := &RouterHandler{
		Engine: *r,
		db:     model.NewDBHandler(),
	}

	v1 := r.Group("/api/v1")
	{
		v1.POST("signup", rh.signupHandler)
		v1.POST("signin", rh.signinHandler)
		post := v1.Group("/post")
		{
			post.GET("img", rh.upLoadImgHandler)
			post.DELETE("img", rh.deleteImgHandler)

			post.POST("content", rh.getPostcontentsHandler)
			post.PATCH("", rh.updatePostHandler)
			post.POST("", rh.upLoadPostHandler) //contents 동시에 가져와야함
		}
	}

	return rh
}
