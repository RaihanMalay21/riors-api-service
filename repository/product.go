package repository

import (
	"github.com/RaihanMalay21/api-service-riors/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	NewTransactionProduct() *gorm.DB
	GetAll() (*[]domain.Product, error)
	Create(tx *gorm.DB, data *domain.Product) error
	UpdateProductImage(tx *gorm.DB, data *domain.Product) error
	GetAllMale() (*[]domain.Product, error)
	GetAllFemale() (*[]domain.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func ConstructorProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (dp *productRepository) NewTransactionProduct() *gorm.DB {
	return dp.db.Begin()
}

func (dp *productRepository) GetAll() (*[]domain.Product, error) {
	var data []domain.Product

	if err := dp.db.Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (dp *productRepository) Create(tx *gorm.DB, data *domain.Product) error {
	if err := tx.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (dp *productRepository) UpdateProductImage(tx *gorm.DB, data *domain.Product) error {
	if err := tx.Model(data).Update("image", data.Image).Error; err != nil {
		return err
	}

	return nil
}

// Get Product by gender

func (dp *productRepository) GetAllMale() (*[]domain.Product, error) {
	var data []domain.Product

	if err := dp.db.Where("category_gender = ?", "Male").Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (dp *productRepository) GetAllFemale() (*[]domain.Product, error) {
	var data []domain.Product

	if err := dp.db.Where("category_gender = ?", "Female").Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
