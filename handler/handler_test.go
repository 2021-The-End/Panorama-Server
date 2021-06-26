package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"panorama/server/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSigninHandler(t *testing.T) {

	router := MakeHandler()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/signin", strings.NewReader(`
	{
		"name":"junwoo",
		"password":"1234",
	}`))
	router.Hh.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, "signin successfully", w.Body.String())

}
func TestSignupHandler(t *testing.T) {

	router := MakeHandler()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/signup", strings.NewReader(`
	{
		"username":"a",
		"password":"124"
	}`))
	router.Hh.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, "{\"msg\":\"signup successfully\",\"statuscode\":201}", w.Body.String())
}

func TestNewPostgre(t *testing.T) {
	dsn := "host=localhost user=postgres password=rlawnsdn6! dbname=panorama port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		panic("failed to connect database")

	}
	log.Println(db)
}

func TestHash(t *testing.T) {
	userpw := "12345"
	hashpw, err := bcrypt.GenerateFromPassword([]byte(userpw), bcrypt.DefaultCost)
	fmt.Println(hashpw, userpw)
	if err != nil {
		panic(err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(userpw))
	fmt.Println(hashpw, userpw)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("ok")
	}
}

func TestVaildation(t *testing.T) {
	router := MakeHandler()

	w := httptest.NewRecorder()

	ck, err := http.Cookie("session_token")
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
	sessionToken := ck
	response := client.Get(sessionToken)
}
