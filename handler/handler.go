package handler

import (
	"errors"
	"log"
	"net/http"
	"panorama/server/httputil"
	"panorama/server/model"
	"panorama/server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

// Signin godoc
// @Summary Signin
// @Description user Signin
// @Tags Users
// @Accept  json
// @Produce  json
// @Router /user/signin [post]
func (rh *RouterHandler) signinHandler(c *gin.Context) {

	var user model.User
	if err = c.ShouldBindJSON(&user); err != nil {
		// If err occurs BINDING(ENCODING) user err, return serverError
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if user.Username == "" {
		// If binded username is empty, return partialcontent
		err = errors.New("username is should'nt be empty")
		httputil.NewError(c, http.StatusPartialContent, err)
		return
	}

	isuser, err := rh.db.SigninIsUser(user)
	if err != nil {
		// If err occurs in calling SignInIsUser, return ISE
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}
	if !isuser {
		// If cant find User, return Unauthorized
		err = errors.New("user not Found")
		httputil.NewError(c, http.StatusUnauthorized, err)
		return
	}
	var httpwriter http.ResponseWriter = c.Writer

	err = utils.GenerateSessionCookie(user.Username, client, httpwriter)

	if err != nil {
		//If err occurs in generating sessioncookie, return ISE
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}
	//Signin successfully
	err = errors.New("login successfully")
	httputil.NewError(c, http.StatusOK, err)

}

// Signup godoc
// @Summary Signup
// @Description user Signup
// @Tags Users
// @Accept  json
// @Produce  json
// @Router /user/signup [post]
func (rh *RouterHandler) signupHandler(c *gin.Context) {
	log.Print("call signup handler")
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		// If err occurs BINDING(ENCODING) user err, return serverError
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	if user.Username == "" {
		// If binded username is empty, return partialcontent
		err = errors.New("username is empty")
		httputil.NewError(c, http.StatusPartialContent, err)
		return
	}
	hashpwd, err := utils.EncrptPasswd(user.Password)
	if err != nil {
		//If err occurs in encrpt passwd, return ISE
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}
	user.Password = hashpwd

	isuser, err := rh.db.SignupIsUser(user)
	if err != nil {
		//If err occurs in calling SignupIsUser, return ISE
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}
	if isuser {
		//If user already exist, return partialcontent
		err = errors.New("user alreay exist")
		httputil.NewError(c, http.StatusPartialContent, err)
		return

	}

	err = rh.db.AddUser(&user)
	if err != nil {
		//If err occurs in Adding User, return partialcontent
		httputil.NewError(c, http.StatusPartialContent, err)
		return
	}
	//Signup successfully
	err = errors.New("signup successfully")
	httputil.NewError(c, http.StatusCreated, err)
}

// uploadImg godoc
// @Summary UploadImg
// @Description Upload Img using fileServer
// @Tags Img
// @Accept  json
// @Produce  json
// @Router /post/img/upload [post]
func (rh *RouterHandler) upLoadImgHandler(c *gin.Context) {

	response, err := utils.Validation(c.Request, client)
	if response == "" {
		httputil.NewError(c, http.StatusUnauthorized, err)
	}
	if err != nil {
		if err == http.ErrNoCookie {
			httputil.NewError(c, http.StatusUnauthorized, err)
			return
		}
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	err = utils.UploadFile(c, response)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
	}
	err = errors.New("successfully Upload")
	httputil.NewError(c, http.StatusOK, err)

}

// DeleteImg godoc
// @Summary DeleteImg
// @Description DeleteImg posted img
// @Tags Img
// @Accept  json
// @Produce  json

// @Router /post/img [delete]
func (rh *RouterHandler) deleteImgHandler(c *gin.Context) {

}

// uploadPost godoc
// @Summary upload Project
// @Description upload Project
// @Tags Project
// @Accept  json
// @Produce  json
// @Router /post [post]
func (rh *RouterHandler) upLoadProjectHandler(c *gin.Context) {
	var post model.Project

	response, err := utils.Validation(c.Request, client)
	if response == "" {
		httputil.NewError(c, http.StatusUnauthorized, err)
	}
	if err != nil {
		if err == http.ErrNoCookie {
			httputil.NewError(c, http.StatusUnauthorized, err)
			return
		}
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}
	if post.ProjectTitle == "" {
		// If binded username is empty, return partialcontent
		err = errors.New("post title empty")
		httputil.NewError(c, http.StatusPartialContent, err)
		return
	}
	if len(post.Contents) < 20 {
		err = errors.New("post contents len should belong then 20")
		httputil.NewError(c, http.StatusPartialContent, err)
		return
	}
	rh.db.UploadPost(&post)
}

// ModifyPost godoc
// @Summary ModifyProject
// @Description ModifyProject
// @Tags Project
// @Accept  json
// @Produce  json
// @Router /post/:id [patch]
func (rh *RouterHandler) modifyProjectHandler(c *gin.Context) {

}

// GetPost godoc
// @Summary GetPost
// @Description Get Post By Id
// @Tags Project
// @Accept  json
// @Produce  json
// @Router /post/:id [get]
func (rh *RouterHandler) getProjectByIdHandler(c *gin.Context) {

	response, err := utils.Validation(c.Request, client)
	if response == "" {
		httputil.NewError(c, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		if err == http.ErrNoCookie {
			httputil.NewError(c, http.StatusUnauthorized, err)
			return
		}
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	id := c.Query("id")
	postid, _ := strconv.Atoi(id)

	post, err := rh.db.GetbyIdPost(postid)

	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}
	log.Println(post)

	// for idx, value := range post.Imgpaths {

	// }
	c.JSON(200, post)

}

func (rh *RouterHandler) reviewHandler(c *gin.Context) {

}

//unresolved yet
// GetEntirePost godoc
// @Summary GetEntirePost
// @Description Get Entire Post to expose introduction page
// @Tags Project
// @Accept  json
// @Produce  json
// @Router /post [get]
func (rh *RouterHandler) getEntireProjectHandler(c *gin.Context) {
	posts, err := rh.db.GetPost()
	if err != nil {
		if posts != nil {
			err = errors.New("empty contents")
			httputil.NewError(c, http.StatusPartialContent, err)
		}
		httputil.NewError(c, http.StatusInternalServerError, err)
	}
	// for i := 0; i < posts[i]/2; i++ {
	// 	posts[i]
	// }

}

//unresolved yet
// uploadComment godoc
// @Summary Comment upload
// @Description Comment upload
// @Tags Comment
// @Accept  json
// @Produce  json
// @Router /comment [post]
func (rh *RouterHandler) uploadCommentHandler(c *gin.Context) {

}

//unresolved yet
// GetComment godoc
// @Summary Get Comment By Id
// @Description Get Comment By postId
// @Tags Comment
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Router /comment/:postid [get]
func (rh *RouterHandler) getCommentHandler(c *gin.Context) {

}

func (rh *RouterHandler) Close() {
	rh.db.Close()
	client.Close()
}
