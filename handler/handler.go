package handler

import (
	"net/http"
	"panorama/server/model"

	"github.com/gin-gonic/gin"
)

// Summary Signin
// Router api/v1/signin [post]
func (rh *RouterHandler) signinHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "binding err" + err.Error()})
		return
	}

	isuser := rh.db.IsUser(&user)
	if !isuser {
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "", "accesstoken": rh.md.GetAccessToken})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "Login Failed : User not Found"})
		return
	}

}

// Summary Sign up
// Router api/v1/signup [post]
func (rh *RouterHandler) signupHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user) //user
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "binding err" + err.Error()})
		return
	}

	hashpwd, err := rh.md.HashPasswd(user.Passwd) //err json
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "hashpwd exchanging err " + err.Error()})
		return
	}

	user.Passwd = hashpwd
	isuser := rh.db.IsUser(&user)

	if !isuser {
		rh.db.AddUser(&user)
	} else {
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "already exist user "})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"statuscode": http.StatusCreated, "msg": "sign up successfully"})
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
