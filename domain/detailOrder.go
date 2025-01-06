package domain

import "gorm.io/gorm"

type DetailOrder struct {
	gorm.Model
	OrderId uint `gorm:"primaryKey"`
	UserId uint `gorm:"bigint;not null"`
	ProductVariantId uint `gorm:"bigint;not null"`
	AmountItems uint `gorm:"bigint;not null"`
	AmountPrice float64 `gorm:"type:DECIMAL(10, 0);not null"`
	Order Order `gorm:"foreignKey:OrderId;references:Id"`
	User User `gorm:"foreignKey:UserId;references:Id"`
	ProductVariant ProductVariant `gorm:"foreignKey:ProductVariantId;references:Id"`
}