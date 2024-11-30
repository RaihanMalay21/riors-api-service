package domain

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id           uint      `gorm:"primaryKey"`
	CreatedAt    time.Time `gorm:"timeStamp"`
	UpdatedAt    time.Time `gorm:"autoCreateTime"`
	CategoryName string    `gorm:"varchar(100);not null;unique"`
	Product      []Product `gorm:"foreignKey:CategoryId"`
}
