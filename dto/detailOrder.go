package dto

type DetailOrder struct {
	OrderId          uint    `json:"orderId"`
	UserId           uint    `json:"userId" validate:"required"`
	ProductVariantId uint    `json:"productVariantId" validate:"required"`
	AmountItems      uint    `json:"amountItems" validate:"required,number"`
	AmountPrice      float64 `json:"amountPrice" validate:"required,number"`
}
