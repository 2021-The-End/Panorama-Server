package model

import "time"

type User struct {
	Name      string    `json:"name"`
	Passwd    string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
type Post struct {
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	Imgpaths  []Image   `json:"imgpaths"` //[{imgpaths:"Asd"},{imgpaths:"asdsa"}]
	CreatedAt time.Time `json:"created_at"`
}
type Image struct {
	ImaPath string
}

type DBHandler interface {
	GetUser() *User
	AddUser(name string) *User
	RemoveUser(id int)
	GetPostContents() *Post
	ModifyPost(*Post)
	Close()
	DeleteImg()
	UploadImg()
}

func NewDBHandler() DBHandler {
	return NewPostgreHandler()
}
