package dto

import "time"

type Order struct {
	Id                   uint               `json:"id"`
	UserId               uint               `json:"userId"`
	TransactionBankId    string             `json:"TransactionBankId"`
	TransactionEwalletId string             `json:"TransactionEwalletId"`
	OrderDate            time.Time          `json:"orderDate"`
	OrderStatus          string             `json:"orderStatus"`
	AmountPrice          float64            `json:"amountPrice"`
	AddressId            uint               `json:"addressId"`
	Address              Address            `json:"Address"`
	User                 User               `json:"user"`
	TransactionBank      TransactionBank    `json:"transactionBank"`
	TransactionEwallet   TransactionEwallet `json:"transactionEwallet"`
	DetailOrder          []DetailOrder
}
