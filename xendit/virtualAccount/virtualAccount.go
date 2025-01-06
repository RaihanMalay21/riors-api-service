package virtualAccount

import (
	"time"

	payment_request "github.com/xendit/xendit-go/v6/payment_request"
)

func CreateChannelPropertiesVA(customerName string, expirationDuration time.Time) payment_request.VirtualAccountChannelProperties {
	return payment_request.VirtualAccountChannelProperties{
		CustomerName: customerName,
		ExpiresAt:    &expirationDuration,
	}
}

func CreateVirtualAccount(amount payment_request.NullableFloat64, currency payment_request.PaymentRequestCurrency, channelCode payment_request.VirtualAccountChannelCode, channelProperties payment_request.VirtualAccountChannelProperties) payment_request.NullableVirtualAccountParameters {
	return *payment_request.NewNullableVirtualAccountParameters(&payment_request.VirtualAccountParameters{
		Amount:            amount,
		Currency:          &currency,
		ChannelCode:       channelCode,
		ChannelProperties: channelProperties,
	})
}

func CreatePaymentMethodVA(typeMethod payment_request.PaymentMethodType, reusabilityMethod payment_request.PaymentMethodReusability, desc payment_request.NullableString, referenceId string, virtualAccount payment_request.NullableVirtualAccountParameters) payment_request.PaymentMethodParameters {
	return payment_request.PaymentMethodParameters{
		Type:           typeMethod,
		Reusability:    reusabilityMethod,
		Description:    desc,
		ReferenceId:    &referenceId,
		VirtualAccount: virtualAccount,
	}
}
