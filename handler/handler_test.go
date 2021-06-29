package handler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"panorama/server/utils"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
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

	username := "junwoo"

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, "signin successfully", w.Body.String())

	SessionKey := uuid.New().String()
	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	log.Println(username)

	err := client.Set(SessionKey, username, 1800*time.Second).Err()
	assert.NoError(t, err)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   SessionKey,
		Expires: time.Now().Add(60 * time.Minute),
	})
	ck, err := req.Cookie("session_token")
	assert.NoError(t, err)

	response, err := utils.Validation(ck.Value, client)
	assert.NoError(t, err)

	assert.NotNil(t, response)

}
func TestSignupHandler(t *testing.T) {

	router := MakeHandler()

	router.db.RemoveUser("a")
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
	log.Println("successfully connect")
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

func TestImgUpload(t *testing.T) {

	router := MakeHandler()

	assert := assert.New(t)
	path := "C:/Users/whktj/Videos/KakaoTalk_20210421_095938464.png"
	file, _ := os.Open(path)
	defer file.Close()

	os.Remove("./public/imgpath")

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "post/img", buf)

	req.Header.Add("Content-type", writer.FormDataContentType())

	router.Hh.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData)
	originFile.Read(originData)
}
