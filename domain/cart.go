package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Id               uint           `gorm:"primaryKey"`
	UserId           uint           `gorm:"bigint;not null"`
	ProductVariantId uint           `gorm:"bigint;not null"`
	AmountItem       float64        `gorm:"type:DECIMAL(10, 0);not null"`
	AmountPrice      float64        `gorm:"type:DECIMAL(10, 0);not null"`
	User             User           `gorm:"foreignKey:UserId;references:Id"`
	ProductVariant   ProductVariant `gorm:"foreignKey:ProductVariantId;references:Id"`
}
