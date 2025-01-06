package users

import (
	"github.com/RaihanMalay21/api-service-riors/domain"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllProduct(userId uint) (*[]domain.Cart, error)
	AddProduct(data *domain.Cart) error
	GetProductVariantId(productId uint, color, size string) (*uint, error)
	UpdateAmountProduct(cart *domain.Cart) error
	DeleteProduct(cart *domain.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func ConstructorCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (cd *cartRepository) GetAllProduct(userId uint) (*[]domain.Cart, error) {
	var data []domain.Cart
	if err := cd.db.Table("riors_cart as rc").
		Where("rc.user_id = ?", userId).
		Joins("JOIN riors_product_variant pv ON pv.id = rc.product_variant_id").
		Joins("JOIN riors_product p ON p.id = pv.product_id").
		Select("rc.*, pv.image, pv.size, pv.color, pv.product_id, p.product_name, p.price").
		Limit(10).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (cd *cartRepository) GetProductVariantId(productId uint, color, size string) (*uint, error) {
	var productVariant domain.ProductVariant
	if err := cd.db.Select("id").Where("product_id = ? AND color = ? AND size = ?", productId, color, size).First(&productVariant).Error; err != nil {
		return nil, err
	}

	return &productVariant.Id, nil
}

func (cd *cartRepository) AddProduct(data *domain.Cart) error {
	if err := cd.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (cd *cartRepository) UpdateAmountProduct(cart *domain.Cart) error {
	if err := cd.db.Model(cart).Updates(map[string]interface{}{
		"amount_price": cart.AmountPrice,
		"amount_item":  cart.AmountItem,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (cd *cartRepository) DeleteProduct(cart *domain.Cart) error {
	if err := cd.db.Where("product_variant_id = ?", cart.ProductVariantId).Delete(cart).Error; err != nil {
		return err
	}

	return nil
}
