package products

import (
	"github.com/RaihanMalay21/api-service-riors/domain"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func ConstructorCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (dc *CategoryRepository) NewTransactionCategory() *gorm.DB {
	return dc.db.Begin()
}

func (dc *CategoryRepository) Create(data *domain.Category) error {
	if err := dc.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (dc *CategoryRepository) GetAll() (*[]domain.Category, error) {
	var data []domain.Category

	if err := dc.db.Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
