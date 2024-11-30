package dto

import "time"

type Category struct {
	Id           uint       `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	CategoryName string     `json:"category" validate:"required,max=100"`
	Product      *[]Product `json:"-" validate:"-"`
}
