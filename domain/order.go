package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id                   uint               `gorm:"primaryKey"`
	UserId               uint               `gorm:"bigint;not null"`
	TransactionBankId    string             `gorm:"vacrhar(200)"`
	TransactionEwalletId string             `gorm:"vacrhar(200)"`
	OrderDate            time.Time          `gorm:"type:date;not null"`
	OrderStatus          string             `gorm:"type:new_order_status"`
	AmountPrice          float64            `gorm:"type:DECIMAL(10, 0)"`
	AddressId            uint               `gorm:"bigint;not null"`
	Address              Address            `gorm:"foreignKey:AddressId;references:Id"`
	User                 User               `gorm:"foreignKey:UserId;references:Id"`
	TransactionBank      TransactionBank    `gorm:"foreignKey:TransactionBankId;references:Id"`
	TransactionEwallet   TransactionEwallet `gorm:"foreignKey:TransactionEwalletId;references:Id"`
	DetailOrder          []DetailOrder
}

// sku
// ekspedi_status
