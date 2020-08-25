package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go_orm/model"
	"log"
	"os"
)

type Repository interface {
	Get(id string) (*model.Person, error)
	Create(id string, name string) error
}

type repo struct {
	DB gorm.DB
}

func CreateRepository(db *gorm.DB) Repository {
	var rep Repository

	rep = &repo{
		*db,
	}
	return rep
}

func (p *repo) Create(id string, name string) error {
	person := &model.Person{
		ID:   id,
		Name: name,
	}
	person.Data = []byte(id + name + "gjfhgjdfgjdjfjdshfjhskjfhkhjdjgjfjdfgjg")

	return p.DB.Create(person).Error
}

func (p *repo) Get(id string) (*model.Person, error) {
	person := new(model.Person)

	err := p.DB.Where("id = ?", id).Find(person).Error

	return person, err
}

func main() {
	// create database gorm_test CHARACTER set  'utf8' collate 'utf8_general_ci';
	db, err := gorm.Open("mysql", "root:QWQ920403@ty@/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Person{})

	// DB, _ := gorm.Open("mysql", db)

	id, name := os.Args[0], os.Args[1]

	repo := CreateRepository(db)
	ret := repo.Create(id, name)
	// ret, err := db.Exec("SELECT * FROM `person` *")
	if ret != nil {
		fmt.Println(ret.Error())
		// t.Error(ret)
	}
	fmt.Println(ret)
}
