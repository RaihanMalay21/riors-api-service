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
	ProductName    string          `gorm:"varchar(100); not null;uniqueProduct" json:"productName"`
	DateRelase     time.Time       `gorm:"type:date;not null" json:"dateRelase"`
	Price          float64         `gorm:"type:DECIMAL(10, 0);not null"  json:"price"`
	Type           string          `gorm:"varchar(100);not null" json:"type"`
	Image          string          `gorm:"varchar(200);not null" json:"image"`
	Desc           string          `gorm:"type:varchar(200);not null" json:"desc"`
	CategoryGender string          `gorm:"type:new_category_gender;not null" json:"categoryGender"`
	Category       Category        `gorm:"foreignKey:CategoryId;references:Id" json:"-"`
	ProductVariant  []ProductVariant `json:"-"`
}
