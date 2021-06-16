package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type MiddleWare interface {
	HashPasswd(string) (string, error)
	CompareHash(string string) bool
	GetAccessToken() string
}

func HashPasswd(userpw string) (string, error) {
	hashpw, err := bcrypt.GenerateFromPassword([]byte(userpw), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return string(hashpw), nil
}

func CompareHash(hashpw, userpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(userpwd))
	if err != nil {
		return false
	} else {
		return true
	}

}

func GetAccessToken() string {

}
