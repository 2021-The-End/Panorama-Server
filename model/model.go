package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"model,omitempty"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

/*
signup 할 때 User에 들어가야 할 것:
Accoutusername,Password

Usersession에 expired까지 들어가야 하는 것
db에 넣을 것인지
in-memory에 넣을 것인지
*/
type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`

	Contents  string    `json:"contents"`
	Imgpaths  []Image   `json:"imgpaths"` //[{imgpaths:"Asd"},{imgpaths:"asdsa"}]
	CreatedAt time.Time `json:"created_at"`
}
type Image struct {
	ImaPath string
}

type DBHandler interface {
	GetUser() *User
	AddUser(*User) error
	RemoveUser(string) error

	UploadPost(*Post) error
	GetPostContents() *Post
	ModifyPost(*Post)

	Close()

	DeleteImg()

	SignupIsUser(User) (bool, error)
	SigninIsUser(User) (bool, error)
}

func NewDBHandler() DBHandler {
	return NewPostgreHandler()
}
