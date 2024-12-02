package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id             uint            `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time       `gorm:"timeStamp" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"autoCreateTime" json:"updated_at"`
	CategoryId     uint            `gorm:"bigint;not null" json:"categoryId"`
	ProductName    string          `gorm:"varchar(100); not null;uniqueProduct" json:"namaProduct"`
	HargaBarang    float64         `gorm:"type:DECIMAL(10, 0);not null"  json:"hargaProduct"`
	Type           string          `gorm:"varchar(100);not null" json:"type"`
	Image          string          `gorm:"varchar(200)" json:"image"`
	CategoryGender string          `gorm:"type:gender_category;not null" json:"categoryGender"`
	Category       Category        `gorm:"foreignKey:CategoryId;references:Id" json:"-"`
	DetailProduct  []DetailProduct `json:"-"`
}
