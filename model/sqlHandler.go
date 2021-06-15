package model

import "github.com/jinzhu/gorm"

type postgreHandler struct {
	db *gorm.DB
}

func NewPostgreHandler() DBHandler {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=panorama sslmode=disable password=rlawnsdn6!")
	if err != nil {

		panic("failed to connect database")

	}
	return &postgreHandler{db: db}
}

func (p *postgreHandler) GetUser() *User {
	return nil
}

func (p *postgreHandler) AddUser(name string) *User {
	p.db.DB().Prepare("UPDATE test1 ")
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
	p.db.Close()
}
func (p *postgreHandler) UploadImg() {

}
