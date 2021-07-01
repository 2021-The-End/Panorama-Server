package model

import (
	"errors"
	"log"
	"panorama/server/utils"
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
	log.Println("call model/SignupIsUser")
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
	log.Println("call model/SigninIsUser")
	result := &User{}

	result.CreatedAt = time.Now()
	result.UpdatedAt = time.Now()
	// Get first matched record
	if err := p.db.Where("username = ?", user.Username).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil //false여야만 singin
		}
		return false, err
	} else {
		if !utils.CompareHash(result.Password, user.Password) || result.Username != user.Username { //db에 있는 username이 json으로 받은 username과 달라야 false 리턴
			err = errors.New("id or Passwd is not match")
			return false, err
		} else {
			return true, nil
		}
	}
}
func (p *postgreHandler) GetUser() *User {
	return nil
}

func (p *postgreHandler) AddUser(user *User) error { //If json.UserName is not equal with db.username AddUser call
	log.Print("call model/AddUser")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := p.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (p *postgreHandler) GetPost() (*[]Post, error) {
	post := &[]Post{}
	if err := p.db.Find(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return post, nil
}

func (p *postgreHandler) RemoveUser(username string) error {
	if err := p.db.Delete(&User{}, username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func (p *postgreHandler) GetbyIdPost(postid int) (*Post, error) {
	result := &Post{}
	if err := p.db.Where("post_id = ?", postid).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}

func (p *postgreHandler) UploadPost(post *Post) error {
	log.Print("call model/UploadPost")
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	if err := p.db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (p *postgreHandler) ModifyPost(*Post) {

}

func (p *postgreHandler) DeleteImg() {

}
func (p *postgreHandler) Close() {
}
