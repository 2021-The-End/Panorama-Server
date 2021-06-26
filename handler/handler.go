package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"panorama/server/model"
	"panorama/server/utils"

	"github.com/gin-gonic/gin"
)

var err error

// Summary Signin
// Router api/v1/signin [post]
func (rh *RouterHandler) signinHandler(c *gin.Context) {
	var user model.User
	if err = c.ShouldBindJSON(&user); err != nil {
		// If err occurs BINDING(ENCODING) user err, return serverError
		utils.ThrowErr(c, http.StatusUnprocessableEntity, err)
		return
	}

	if user.Username == "" {
		// If binded username is empty, return partialcontent
		err = errors.New("username is should'nt be empty")
		utils.ThrowErr(c, http.StatusPartialContent, err)
		return
	}

	isuser, err := rh.db.SigninIsUser(user)
	if err != nil {
		// If err occurs in calling SignInIsUser, return ISE
		utils.ThrowErr(c, http.StatusPartialContent, err)
		return
	}
	if !isuser {
		// If cant find User, return Unauthorized?
		err = errors.New("user not Found")
		utils.ThrowErr(c, http.StatusUnauthorized, err)
		return
	}
	err = utils.GenerateSessionCookie(user.Username, client, c)

	if err != nil {
		//If err occurs in generating sessioncookie, return ISE
		utils.ThrowErr(c, http.StatusInternalServerError, err)
		return
	}
	//Signin successfully
	err = errors.New("login successfully")
	utils.ThrowErr(c, http.StatusOK, err)
}

// Summary Sign up
// Router api/v1/signup [post]
func (rh *RouterHandler) signupHandler(c *gin.Context) {
	log.Print("call signup handler")
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		// If err occurs BINDING(ENCODING) user err, return serverError
		utils.ThrowErr(c, http.StatusInternalServerError, err)
		return
	}
	if user.Username == "" {
		// If binded username is empty, return partialcontent
		err = errors.New("username is empty")
		utils.ThrowErr(c, http.StatusPartialContent, err)
		return
	}
	hashpwd, err := utils.EncrptPasswd(user.Password)
	if err != nil {
		//If err occurs in encrpt passwd, return ISE
		utils.ThrowErr(c, http.StatusInternalServerError, err)
		return
	}
	user.Password = hashpwd

	isuser, err := rh.db.SignupIsUser(user)
	if err != nil {
		//If err occurs in calling SignupIsUser, return ISE
		utils.ThrowErr(c, http.StatusInternalServerError, err)
		return
	}
	if isuser {
		//If user already exist, return partialcontent
		err = errors.New("user alreay exist")
		utils.ThrowErr(c, http.StatusPartialContent, err)
		return

	}

	err = rh.db.AddUser(&user)
	if err != nil {
		//If err occurs in Adding User, return partialcontent
		utils.ThrowErr(c, http.StatusPartialContent, err)
		return
	}
	//Signup successfully
	err = errors.New("signup successfully")
	utils.ThrowErr(c, http.StatusCreated, err)
}

// Summary upload img
// Description Upload img to public folder to use fileserver
// Router api/v1/post/img [get]
func (rh *RouterHandler) upLoadImgHandler(c *gin.Context) { //cookie
	ck, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			utils.ThrowErr(c, http.StatusUnauthorized, err)
			return
		}
		// For any other type of error, return a bad request status
		utils.ThrowErr(c, http.StatusInternalServerError, err)
		return
	}

	response, err := utils.Validation(ck, client)
	if err != nil {
		utils.ThrowErr(c, http.StatusUnauthorized, err)
		return
	}
	response.Name()
	//access token 검증-> 그 key인 username 가져오기
	header, err := c.FormFile("upload_file")
	uploadfile, _ := header.Open()
	defer uploadfile.Close()
	if err != nil {
		utils.ThrowErr(c, http.StatusInternalServerError, err)
	}

	dirname := "./imgpath/"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", "uploads", header.Filename)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		fmt.Fprint(c.Writer, err)
		return
	}
	io.Copy(file, uploadfile)
	c.Status(200)
	fmt.Fprint(c.Writer, filepath)
	c.Redirect(http.StatusTemporaryRedirect, "/public/index.html")

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
	client.Close()
}
