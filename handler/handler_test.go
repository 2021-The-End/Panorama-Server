package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
