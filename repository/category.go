package repository

import (
	"github.com/RaihanMalay21/api-service-riors/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	NewTransactionCategory() *gorm.DB
	Create(data *domain.Category) error
	GetAll() (*[]domain.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func ConstructorCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (dc *categoryRepository) NewTransactionCategory() *gorm.DB {
	return dc.db.Begin()
}

func (dc *categoryRepository) Create(data *domain.Category) error {
	if err := dc.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (dc *categoryRepository) GetAll() (*[]domain.Category, error) {
	var data []domain.Category

	if err := dc.db.Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
