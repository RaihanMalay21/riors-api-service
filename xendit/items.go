package xendit

import "github.com/xendit/xendit-go/v6/payment_request"

func CreateBasketItem(productId string, productName string, quantity float64, price float64, category string, currency string, productType string) payment_request.PaymentRequestBasketItem {
	return payment_request.PaymentRequestBasketItem{
		ReferenceId: &productId,
		Name:        productName,
		Quantity:    quantity,
		Price:       price,
		Category:    category,
		Currency:    currency,
		Type:        &productType,
	}
}
