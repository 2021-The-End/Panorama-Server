package handler

import (
	"panorama/server/model"

	"github.com/gin-gonic/gin"
)

type Success struct {
	success bool
}

// Summary Signin
// Router api/v1/signin [post]
func (rh *RouterHandler) signinHandler(c *gin.Context) {
	var user model.User
	c.Bind(&user)
}

// Summary Sign up
// Router api/v1/signup [post]
func (rh *RouterHandler) signupHandler(c *gin.Context) {
	var user model.User
	c.Bind(&user) //user
	rh.md.HashPasswd(user.Passwd)
	isuser := rh.md.IsUser(&user)

	if isuser {
		rh.db.AddUser(user.Name)
	}

}

// Summary upload img
// Description Upload img to public folder to use fileserver
// Router api/v1/post/img [get]
func (rh *RouterHandler) upLoadImgHandler(c *gin.Context) {

}

// Summary get post contents
// Router api/v1/post/content [post]
func (rh *RouterHandler) getPostcontentsHandler(c *gin.Context) {

}

// Summary upload post
// Router api/v1/post [post]
func (rh *RouterHandler) upLoadPostHandler(c *gin.Context) {

}

// Summary update post contents
// Router api/v1/post [patch]
func (rh *RouterHandler) updatePostHandler(c *gin.Context) {

}

// Summary delete img temporary
// Router api/v1/img [delete]
func (rh *RouterHandler) deleteImgHandler(c *gin.Context) {

}

func (rh *RouterHandler) Close() {
	rh.db.Close()
}
