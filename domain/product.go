package domain

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id uint	`gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"timeStamp"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	CategoryId uint `gorm:"bigint;not null"`
	ProductName string `gorm:"varchar(100); not null;uniqueProduct"`
	HargaBarang float64 `gorm:"type:DECIMAL(10, 0);not null"`
	Type string `gorm:"varchar(100);not null"`
	Image string `gorm:"varchar(200)"`
	Category Category `gorm:"foreignKey:CategoryId;references:Id"`
	DetailProduct []DetailProduct
}