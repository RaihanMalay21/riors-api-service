package ewallet

import "github.com/xendit/xendit-go/v6/payment_request"

func CreateChannelPropertiesEwallet(succesReturnUrl, failureReturnUrl, cancelReturnUrl, mobileNumber *string) payment_request.EWalletChannelProperties {
	return payment_request.EWalletChannelProperties{
		SuccessReturnUrl: succesReturnUrl,
		FailureReturnUrl: failureReturnUrl,
		CancelReturnUrl:  cancelReturnUrl,
		MobileNumber:     mobileNumber,
	}
}

func CreateEwallet(ewallet payment_request.EWalletChannelCode, channelProperties payment_request.EWalletChannelProperties) payment_request.NullableEWalletParameters {
	return *payment_request.NewNullableEWalletParameters(&payment_request.EWalletParameters{
		ChannelCode:       &ewallet,
		ChannelProperties: &channelProperties,
	})
}

func CreatePaymentMethodEWALLET(typeMethod payment_request.PaymentMethodType, reusabilityMethod payment_request.PaymentMethodReusability, desc payment_request.NullableString, referenceId string, ewallet payment_request.NullableEWalletParameters) payment_request.PaymentMethodParameters {
	return payment_request.PaymentMethodParameters{
		Type:        typeMethod,
		Reusability: reusabilityMethod,
		Description: desc,
		ReferenceId: &referenceId,
		Ewallet:     ewallet,
	}
}
