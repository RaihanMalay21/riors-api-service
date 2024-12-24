package dto

import (
	"time"
)

type User struct {
	Id uint `json:"id"` 
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"json:"UpdatedAt"` 
	Name string `validate:"required,min=6,max=100"`
	Email string `validate:"-" json:"email"`
	Whatshapp int `validate:"uniqueWA,required,number"`
	Password string `validate:"-"`
	Poin float64 `json:"poin"`
	Address []Address `gorm:"foreignKey:UserID" validate:"-" json:"address"`
}

