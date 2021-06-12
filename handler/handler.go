package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Fprint(c.Writer, "index")
}

// getUsersListHandler godoc
// @Summary Get User List
// @Description Get entire Users Info
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200
// @Router /users [get]
func getUsersListHandler(c *gin.Context) {

}

//getUserHandler
// @Summary Get User Info
// @Accept  json
// @Produce  json
// @Success 200
// @Description Get entire Users Info
// @Router /users/{id} [get]
func getUserHandler(c *gin.Context) {

}

// createUsersHandler
// @Summary Create User
// @Description Get entire Users Info
// @Accept  json
// @Produce  json
// @Success 200
// @Router /users [post]
func createUsersHandler(c *gin.Context) {

}

// updateUsersHandler
// @Summary Create User
// @Description Get entire Users Info
// @Accept  json
// @Produce  json
// @Success 200
// @Router /users/{id} [get]
func updateUsersHandler(c *gin.Context) {

}
