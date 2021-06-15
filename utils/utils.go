package utils

import (
	"panorama/server/model"

	"golang.org/x/crypto/bcrypt"
)

type MiddleWare interface {
	IsUser(*model.User) bool
	HashPasswd(string) (string, ok bool)
	CompareHash(string string) bool
}

func HashPasswd(userpw string) string {
	hashpw, _ := bcrypt.GenerateFromPassword([]byte(userpw), bcrypt.DefaultCost)
	return string(hashpw)
}

func CompareHash(hashpw, userpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(userpwd))
	if err != nil {
		return false
	} else {
		return true
	}

}
