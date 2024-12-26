package dto

import (
	"time"
)

type User struct {
	Id uint `json:"id"` 
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"` 
	Name string `validate:"required,min=6,max=100"`
	Email string `validate:"-" json:"email"`
	Whatsapp string `validate:"uniqueWAUser,required,whatsapp"`
	Password string `validate:"-"`
	Poin float64 `json:"poin"`
	Address []Address `validate:"-" json:"address"`
}

