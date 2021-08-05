package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"panorama/server/utils"
	"time"

	"panorama/server/info"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgreHandler struct {
	sqldb *sql.DB
	ormdb *gorm.DB
}

var dsn = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	info.DBHost, info.DBPort, info.User, info.Password, info.Dbname)

func NewPostgreHandler() DBHandler {
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect sql database")
	}
	ormdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("failed to connect orm database")

	}
	return &postgreHandler{sqldb: sqldb, ormdb: ormdb}
}

//if value is not user, return false
func (p *postgreHandler) SignupIsUser(user User) (bool, error) {
	log.Println("call model/SignupIsUser")
	result := user
	// Get first matched record
	if err := p.ormdb.Where("username = ?", result.Username).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil //false여야만 signup
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

	// Get first matched record
	if err := p.ormdb.Where("username = ?", user.Username).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
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

	if err := p.ormdb.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (p *postgreHandler) GetPost() (*[]ProjectSum, error) {
	post := &[]ProjectSum{}

	if err := p.ormdb.Find(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return post, nil
}

func (p *postgreHandler) RemoveUser(username string) error {
	if err := p.ormdb.Delete(&User{}, username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func (p *postgreHandler) GetbyIdPost(projectid int) (*Projectcon, error) {
	Projectcon := &Projectcon{}
	var img []string
	var cre []string
	selImgQuery := info.SelImgpathQuery
	selCreQuery := info.SelCreaterQuery

	err := p.sqldb.QueryRow(selImgQuery, projectid).Scan(pq.Array(&img))
	if err != nil {
		return nil, err
	}

	Projectcon.Imgpaths = img
	log.Println(img)

	err = p.sqldb.QueryRow(selCreQuery, projectid).Scan(pq.Array(&cre))
	if err != nil {
		return nil, err
	}
	Projectcon.Creaters = cre
	log.Println(cre)

	if err := p.ormdb.Omit("Imgpaths", "Creaters").Table("projectcon").First(&Projectcon, "id = ?", projectid).Error; err != nil {
		return nil, err
	}
	return Projectcon, nil
}

func (p *postgreHandler) UploadPost(Projectcon *Projectcon) error {
	log.Print("call model/UploadPost")
	Projectcon.CreatedAt = time.Now()
	Projectcon.UpdatedAt = time.Now()

	insQuery := info.InsQuery
	_, err := p.sqldb.Exec(insQuery, Projectcon.Title, Projectcon.Contents,
		pq.Array(Projectcon.Creaters), pq.Array(Projectcon.Imgpaths), Projectcon.Summary, Projectcon.Grade, Projectcon.CreatedAt, Projectcon.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (p *postgreHandler) ModifyPost(Projectcon *JSONProjectCon) error {
	log.Print("call model/ModifyPost")
	Projectcon.CreatedAt = time.Now()
	Projectcon.UpdatedAt = time.Now()
	if err := p.ormdb.Model(&Projectcon).Where("id = ?", Projectcon.ID).Updates(&Projectcon).Error; err != nil {
		return err
	}
	return nil
}

func (p *postgreHandler) DeleteImg() {

}
func (p *postgreHandler) Close() {
	p.sqldb.Close()
}
