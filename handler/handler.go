package handler

import (
	"log"
	"net/http"
	"panorama/server/model"

	"panorama/server/utils"

	"github.com/gin-gonic/gin"
)

// Summary Signin
// Router api/v1/signin [post]
func (rh *RouterHandler) signinHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"statuscode": http.StatusUnprocessableEntity, "msg": "Invalid json provided" + err.Error()})
		return
	}

	if user.Username == "" {
		log.Println("username is empty")
		c.JSON(http.StatusPartialContent, gin.H{"statuscode": http.StatusPartialContent, "msg": "username should be not null"})
		return
	}

	isuser, err := rh.db.SigninIsUser(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"statuscode": http.StatusInternalServerError, "msg": err.Error()})
	}
	if !isuser {
		sessionkey, err := utils.GenerateSessionCookie(user, cache, c)

		if err != nil {
			log.Println("generate Cookie err")
			c.JSON(http.StatusPartialContent, gin.H{"statuscode": http.StatusInternalServerError, "msg": err.Error()})
			return
		}

		log.Println("successful signin")
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "signin successfully", "accesstoken": sessionkey})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"statuscode": http.StatusUnauthorized, "msg": "Login Failed : User not Found"})
		return
	}
}

// Summary Sign up
// Router api/v1/signup [post]
func (rh *RouterHandler) signupHandler(c *gin.Context) {
	log.Print("call signup handler")
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "binding err " + err.Error()})
		return
	}
	if user.Username == "" {
		log.Println("username is empty")
		c.JSON(http.StatusPartialContent, gin.H{"statuscode": http.StatusPartialContent, "msg": "username should be not null"})
	}
	hashpwd, err := utils.EncrptPasswd(user.Password) //err json
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "hashpwd exchanging err " + err.Error()})
		return
	}

	user.Password = hashpwd
	isuser, err := rh.db.SignupIsUser(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": err.Error()})
		return
	}
	if isuser {
		log.Println("user alreay exist")
		c.JSON(http.StatusOK, gin.H{"statuscode": http.StatusOK, "msg": "already exist user"})
		return

	} else {
		err = rh.db.AddUser(&user)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusCreated, gin.H{"statuscode": http.StatusCreated, "msg": "unable to add user"})
			return
		} else {
			log.Println("successful signup")
			c.JSON(http.StatusCreated, gin.H{"statuscode": http.StatusCreated, "msg": "signup successfully"})
			return
		}
	}
}

// Summary upload img
// Description Upload img to public folder to use fileserver
// Router api/v1/post/img [get]
func (rh *RouterHandler) upLoadImgHandler(c *gin.Context) {
	ck, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			c.JSON(http.StatusUnauthorized, gin.H{"statuscode": http.StatusUnauthorized, "msg": "Cookie is Unavailable"})
			return
		}
		// For any other type of error, return a bad request status
		c.JSON(http.StatusBadRequest, gin.H{"statuscode": http.StatusBadRequest, "msg": "caused err to load cookie"})
		return
	}
	sessionToken := ck
	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		c.JSON(http.StatusInternalServerError, gin.H{"statuscode": http.StatusInternalServerError, "msg": "caused cache err"})
		return
	}
	if response == nil {
		// If the session token is not present in cache, return an unauthorized error
		c.JSON(http.StatusUnauthorized, gin.H{"statuscode": http.StatusUnauthorized})
		return
	}
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
