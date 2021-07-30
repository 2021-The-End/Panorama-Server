package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Username string `json:"username"`
	Password string `json:"password"`
}

//ProjectSum is PostSummary to get "전시소개" Summary of entire post
type ProjectSum struct {
	*gorm.Model

	Title    string   `json:"title"`
	Creaters []string `json:"creaters"`
	Grade    int      `json:"grade"`
}

//json Project
type JSONProjectCon struct {
	*gorm.Model
	Title    string   `json:"title"`
	Creaters []string `json:"creaters"`
	Grade    int      `json:"grade"`

	Summary  string   `json:"summary"`
	Contents string   `json:"contents"`
	Imgpaths []string `json:"imgpaths"`
}
type Projectcon struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title    string   `json:"title"`
	Creaters []string `json:"creaters"`
	Grade    int      `json:"grade"`

	Summary  string   `json:"summary"`
	Contents string   `json:"contents"`
	Imgpaths []string `json:"imgpaths"`
}

type Comment struct {
	ProjectId int    `json:"project_id"`
	Contents  string `json:"contents"`
}
type DBHandler interface {
	GetUser() *User
	AddUser(*User) error
	RemoveUser(string) error

	UploadPost(*Projectcon) error
	GetbyIdPost(postid int) (*Projectcon, error)
	ModifyPost(ProjectCon *JSONProjectCon) error
	GetPost() (*[]ProjectSum, error)

	Close()

	DeleteImg()

	SignupIsUser(User) (bool, error)
	SigninIsUser(User) (bool, error)
}

func NewDBHandler() DBHandler {
	return NewPostgreHandler()
}
