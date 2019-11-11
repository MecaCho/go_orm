package go_orm

import (
	"go_orm/model"
	"github.com/jinzhu/gorm"
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

	return p.DB.Create(person).Error
}

func (p *repo) Get(id string) (*model.Person, error) {
	person := new(model.Person)

	err := p.DB.Where("id = ?", id).Find(person).Error

	return person, err
}
