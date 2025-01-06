package dto

type Cart struct {
	Id               uint    `json:"id"`
	ProductId        uint    `json:"productId" validate:"required"` // add product or update
	UserId           uint    `json:"userId" validate:"required"`
	Color            string  `json:"color" validate:"required"`               // add product or update
	Size             string  `json:"size" validate:"required"`                // add product or update
	AmountItem       float64 `json:"amountItem" validate:"required,number"`   // add product or update
	AmountPrice      float64 `json:"amountPrice" validate:"required, number"` // add prodcut or update
	ProductVariantId uint    `json:"productVariantId"`
	Image            string  `json:"image"`
	ProductName      string  `json:"productName"`
	Price            float64 `json:"price"`
}
