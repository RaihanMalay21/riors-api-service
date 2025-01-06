package dto

import (
	"time"
)

type TransactionBank struct {
	Id                   string    `json:"id"`
	OrderId              uint      `json:"orderId"`
	StatusTransaction    string    `json:"statusTransaction"`
	VirtualAccountNumber string    `json:"virtualAccountNumber"`
	Amount               float64   `json:"amount"`
	ExpiresAt            time.Time `json:"expiresAt"`
	ChannelCode          string    `json:"channelCode"`
	PaidAt               time.Time `json:"paidAt"`
}

type TransactionEwallet struct {
	Id                string    `json:"id"`
	OrderId           uint      `json:"orderId"`
	StatusTransaction string    `json:"statusTransaction"`
	Amount            float64   `json:"amount"`
	MobileNumber      string    `json:"mobileNumber"`
	ChannelCode       string    `json:"channelCode"`
	ExpiresAt         time.Time `json:"expiresAt"`
}
