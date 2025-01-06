package products

import (
	"github.com/RaihanMalay21/api-service-riors/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func ConstructorProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (dp *ProductRepository) NewTransactionProduct() *gorm.DB {
	return dp.db.Begin()
}

func (dp *ProductRepository) GetAll() (*[]domain.Product, error) {
	var data []domain.Product

	if err := dp.db.Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (dp *ProductRepository) Create(tx *gorm.DB, data *domain.Product) error {
	if err := tx.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (dp *ProductRepository) UpdateProductImage(tx *gorm.DB, data *domain.Product) error {
	if err := tx.Model(data).Update("image", data.Image).Error; err != nil {
		return err
	}

	return nil
}

// Get Product by gender

func (dp *ProductRepository) GetAllMale() (*[]domain.Product, error) {
	var data []domain.Product

	if err := dp.db.Where("category_gender = ?", "Man").Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (dp *ProductRepository) GetAllFemale() (*[]domain.Product, error) {
	var data []domain.Product

	if err := dp.db.Where("category_gender = ?", "Woman").Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
