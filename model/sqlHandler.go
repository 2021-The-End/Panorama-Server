package model

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgreHandler struct {
	db *gorm.DB
}

func NewPostgreHandler() DBHandler {
	dsn := "host=localhost user=postgres password=rlawnsdn6! dbname=panorama port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		panic("failed to connect database")

	}
	return &postgreHandler{db: db}
}

//if value is not user, return false
func (p *postgreHandler) SignupIsUser(user User) (bool, error) {
	log.Println("call model/IsUser")
	result := user

	result.CreatedAt = time.Now()
	result.UpdatedAt = time.Now()
	// Get first matched record
	if err := p.db.Where("username = ?", result.Username).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return true, err
	} else {
		if result.Username != user.Username { //db에 있는 username이 json으로 받은 username과 달라야 false 리턴
			return false, nil
		} else {
			return true, nil
		}
	}
}
func (p *postgreHandler) SigninIsUser(user User) (bool, error) {
	return false, nil
}
func (p *postgreHandler) GetUser() *User {
	return nil
}

func (p *postgreHandler) AddUser(user *User) error { //If json.UserName is not equal with db.username AddUser call
	log.Print("call model/AddUser")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := p.db.Omit("deleted_at").Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (p *postgreHandler) RemoveUser(id int) {

}

func (p *postgreHandler) GetPostContents() *Post {
	return nil
}
func (p *postgreHandler) ModifyPost(*Post) {

}

func (p *postgreHandler) DeleteImg() {

}
func (p *postgreHandler) Close() {
}
func (p *postgreHandler) UploadImg() {

}
