package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint      `gorm:"primaryKey" query:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"UpdatedAt"`
	UserName  string    `gorm:"varchar(100);unique"`
	Email     string    `gorm:"varchar(100);unique"`
	Whatshapp int       `gorm:"int;unique"`
	Password  string    `gorm:"varchar(200)"`
	Poin      float64   `gorm:"type:DECIMAL(10, 0)"`
	Address   []Address `gorm:"foreignKey:UserId"`
}
