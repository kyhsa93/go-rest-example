package repository

import (
	"study/dto"
	"study/model"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) Save(data *dto.Study) {
	study := &model.Study{Name: data.Name, Description: data.Description}
	err := repository.db.Save(study).Error
	if err != nil {
		panic(err)
	}
}

func (repository *Repository) FindAll() model.Studies {
	var studies model.Studies
	err := repository.db.Find(&studies).Error
	if err != nil {
		panic(err)
	}
	return studies
}

func (repository *Repository) FindById(id string) model.Study {
	var study model.Study
	repository.db.Where(&model.Study{Model: model.Model{ID: id}}).Take(&study)
	return study
}

func (repository *Repository) Delete(id string) {
	err := repository.db.Delete(&model.Study{Model: model.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}