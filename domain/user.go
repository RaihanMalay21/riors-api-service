package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id         uint      `gorm:"primaryKey" json:"id" query:"id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt  time.Time `gorm:"autoCreateTime" json:"UpdatedAt"`
	Name       string    `gorm:"varchar(100)" json:"name"`
	Email      string    `gorm:"varchar(100);unique;not null" json:"email"`
	Whatsapp   string    `gorm:"varchar(20);unique" json:"whatsapp"`
	Password   string    `gorm:"varchar(200)" json:"-"`
	Poin       float64   `gorm:"type:DECIMAL(10, 0)" json:"poin"`
	Active     bool      `gorm:"boolean"`
	LastActive time.Time `gorm:"timestamp"`
	Address    []Address `gorm:"foreignKey:UserId"`
}
