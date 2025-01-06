package dto

import (
	"time"
)

type Product struct {
	Id             uint             `json:"id"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	CategoryId     uint             `json:"categoryId" validate:"required"`
	ProductName    string           `json:"productName" validate:"required,max=100,uniqueProduct"`
	DateRelase     string           `json:"dateRelase" validate:"required,date_format"`
	Price          float64          `json:"price" validate:"required,number"`
	Desc           string           `json:"desc" validate:"required,max=200"`
	Type           string           `json:"type" validate:"required,max=100"`
	Image          string           `json:"image" validate:"required"`
	CategoryGender string           `json:"categoryGender" validate:"required"`
	FileSize       uint             `json:"-" validate:"required,maxSizeFile"`
	Ext            string           `json:"-" validate:"required"`
	ImageType      string           `json:"-"  validate:"required"`
	Category       Category         `json:"-" validate:"-"`
	DetailProduct  *[]DetailProduct `json:"-" validate:"-"`
	// Ext string  validate:"typeExt"
}
