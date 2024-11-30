package dto

import (
	"time"
)

type DetailProduct struct {
	Id uint `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductId uint `json:"productId" validate:"required"`
	Size string `json:"size" validate:"required,max=20"`
	Color string `json:"color" validate:"required,max=50"`
	Stock uint `json:"stock" validate:"required,number"`
	Product Product `json:"product" validate:"-"`
}