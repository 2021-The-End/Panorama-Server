package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func EncrptPasswd(userpw string) (string, error) {
	hashpw, err := bcrypt.GenerateFromPassword([]byte(userpw), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return string(hashpw), nil
}

func CompareHash(hashpw, userpw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(userpw))
	if err != nil {
		return false
	} else {
		return true
	}

}

func IsexistAccessToken() bool {
	return false
}

func GenerateSessionCookie(username string, client *redis.Client, c *gin.Context) error {
	SessionKey := uuid.New().String()
	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	log.Println(username)

	err := client.Set(SessionKey, username, 1800*time.Second).Err()

	if err != nil {
		return err
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   SessionKey,
		Expires: time.Now().Add(60 * time.Minute),
	})
	return nil
}

func ThrowErr(c *gin.Context, statuscode int, err error) {
	log.Println(err)
	c.JSON(statuscode, gin.H{"statuscode": statuscode, "msg": err.Error()})
}

func Validation(sessionToken string, client *redis.Client) (string, error) {

	response, err := client.Get(sessionToken).Result()
	if response == "" {
		// If the session token is not present in cache, return an unauthorized error
		err := errors.New("session token is not present in cache")
		return "", err
	}
	if err != nil {
		return "", err
	}

	return response, nil
}

func UploadFile(c *gin.Context, response string) error {
	header, err := c.FormFile("upload_file")
	uploadfile, _ := header.Open()
	if err != nil {
		return err
	}
	defer uploadfile.Close()

	dirname := "./public/imgpath/" + response
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s/%s", "public/imgpath", response, header.Filename) //imgpath/username/filename
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, uploadfile)

	return nil
}
