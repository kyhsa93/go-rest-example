package repository

import (
	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/entity"
	"github.com/kyhsa93/gin-rest-example/config"

	"github.com/jinzhu/gorm"
)

// Repository repository for query to database
type Repository struct {
	database *gorm.DB
}

// NewRepository create repository instance
func NewRepository(config *config.Config) *Repository {
	database := config.Database.Connection
	database.AutoMigrate(&entity.Account{})
	return &Repository{database: database}
}

func (repository *Repository) dtoToEntity(dto *dto.Account) *entity.Account {
	return &entity.Account{Email: dto.Email, SocialID: dto.SocialID}
}

// Save create or update account
func (repository *Repository) Save(data *dto.Account, accountID string) {
	accountEntity := repository.dtoToEntity(data)

	if accountID != "" {
		accountEntity.ID = accountID
	}

	err := repository.database.Save(accountEntity).Error

	if err != nil {
		panic(err)
	}
}

// FindByEmailAndSocialID find all account
func (repository *Repository) FindByEmailAndSocialID(email string, socialID string) (accountEntity entity.Account) {
	repository.database.Where(&entity.Account{Email: email, SocialID: socialID}).First(&accountEntity)
	return
}

// FindByID find account by accountId
func (repository *Repository) FindByID(id string) (accountEntity entity.Account) {
	repository.database.Where(&entity.Account{Model: entity.Model{ID: id}}).First(&accountEntity)
	return
}

// Delete delete account by accountId
func (repository *Repository) Delete(id string) {
	err := repository.database.Delete(&entity.Account{Model: entity.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}
