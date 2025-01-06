package domain

import "gorm.io/gorm"

type ProductVariant struct {
	gorm.Model
	Id        uint    `gorm:"primaryKey"`
	ProductId uint    `gorm:"bigint"`
	Color     string  `gorm:"varchar(50);not null"`
	Size      string  `gorm:"varchar(50);not null"`
	Stock     uint64  `gorm:"bigint"`
	Image     string  `gorm:"varchar(200);not null"`
	Product   Product `gorm:"foreignKey:productId;references:Id"`
}
