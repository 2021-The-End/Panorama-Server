package utils

import (
	"log"
	"net/http"
	"panorama/server/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
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

func GenerateSessionCookie(user model.User, cache redis.Conn, c *gin.Context) (string, error) {
	var result model.UserSession
	result.User = user

	result.SessionKey = uuid.New().String()
	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	_, err := cache.Do("SETEX", result.SessionKey, "120", result.User.Username)
	if err != nil {
		// If there is an error in setting the cache, return an internal server error
		return "", err
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   result.SessionKey,
		Expires: time.Now().Add(120 * time.Second),
	})
	return result.SessionKey, nil
}

func IsexistAccessToken() bool {
	return false
}
