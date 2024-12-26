package domain

import (

)

type Address struct {
	Id uint `gorm:"primaryKey"`
	UserId uint `gorm:"bigint;not null"`
	KodePos uint `gorm:"bigint;not null"`
	AcceptName string `gorm:"varchar(100);not null"`
	Province string `gorm:"varchar(50);not null"`
	KabupatenKota string `gorm:"varchar(100);not null"`
	Kecamatan string `gorm:"varchar(100);not null"`
	Desa string `gorm:"varchar(100);not null"`
	RT_RW string `gorm:"varchar(20);not null"`
	NoHouse int `gorm:"int;not null"`
	StreetName string `gorm:"varchar(100)"`
	DetailAddress string `gorm:"Text;not null"`
	User User `gorm:"foreignKey:UserId;references:Id"`
}