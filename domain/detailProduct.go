package domain

import (
	"gorm.io/gorm"
	"time"
)

type DetailProduct struct {
	gorm.Model
	Id uint `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"timeStamp"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	ProductId uint `gorm:"bigint;not null"`
	Size string `gorm:"varchar(20);not null"`
	Color string `gorm:"varchar(50);not null"`
	Stock uint `gorm:"bigint;default(0)"`
	Product Product `gorm:"foreignkey:ProductId;references:Id"`
}