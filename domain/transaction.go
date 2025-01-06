package domain

import (
	"time"

	"gorm.io/gorm"
)

type TransactionBank struct {
	gorm.Model
	Id                   string    `gorm:"primaryKey"`
	OrderId              uint      `gorm:"bigint;not null"`
	StatusTransaction    string    `gorm:"varchar(50);not null"`
	VirtualAccountNumber string    `gorm:"varchar(50);not null"`
	Amount               float64   `gorm:"type:DECIMAL(10, 0);not null"`
	ExpiresAt            time.Time `gorm:"timestamp;not null"`
	ChannelCode          string    `gorm:"varchar(20);not null"`
	PaidAt               time.Time `gorm:"timestamp;not null"`
}

type TransactionEwallet struct {
	gorm.Model
	Id                string    `gorm:"primaryKey"`
	OrderId           uint      `gorm:"bigint;not null"`
	StatusTransaction string    `gorm:"varchar(50);not null"`
	Amount            float64   `gorm:"type:DECIMAL(10, 0);not null"`
	MobileNumber      string    `gorm:"varchar(50);not null"`
	ChannelCode       string    `gorm:"vacrhar(20);not null"`
	ExpiresAt         time.Time `gorm:"timestamp;not null"`
}
