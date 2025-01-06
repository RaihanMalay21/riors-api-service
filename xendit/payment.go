package xendit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/xendit/xendit-go/v6"
	payment_request "github.com/xendit/xendit-go/v6/payment_request"
)

type PaymentXendit struct {
	xenditClient *xendit.APIClient
}

func (p *PaymentXendit) XenditClient() {
	client := xendit.NewClient("xnd_development_8jTP7KPfzhIQqRdwebLJYPYO4rru7R904RL8Kxg96nhLBQSwkknvZcvnEdNXk1")
	p.xenditClient = client
}

func (p *PaymentXendit) CreatePaymentRequestVA(
	currency payment_request.PaymentRequestCurrency,
	referenceId string,
	amountPrice float64,
	desc payment_request.NullableString,
	paymentMethod payment_request.PaymentMethodParameters,
	// customer map[string]interface{},
	customerId payment_request.NullableString,
) (map[string]interface{}, error) {
	paymentRequestParameter := payment_request.PaymentRequestParameters{
		Currency:      currency,
		ReferenceId:   &referenceId,
		PaymentMethod: &paymentMethod,
		Description:   desc,
		Amount:        &amountPrice,
		CustomerId:    customerId,
	}

	res, _, err := p.xenditClient.PaymentRequestApi.CreatePaymentRequest(context.Background()).
		PaymentRequestParameters(paymentRequestParameter).
		Execute()
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"Id":                   res.Id,
		"Created":              res.Created,
		"ReferenceId":          res.ReferenceId,
		"BusinessId":           res.BusinessId,
		"CustemerId":           *res.CustomerId.Get(),
		"Amount":               fmt.Sprintf("%.0f", *res.PaymentMethod.VirtualAccount.Get().Amount.Get()),
		"Status":               res.Status,
		"ExpiresAt":            *res.PaymentMethod.VirtualAccount.Get().ChannelProperties.ExpiresAt,
		"Type":                 res.PaymentMethod.Type,
		"VirtualAccountNumber": *res.PaymentMethod.VirtualAccount.Get().ChannelProperties.VirtualAccountNumber,
		"ChannelCode":          res.PaymentMethod.VirtualAccount.Get().ChannelCode,
	}

	return response, nil
}

func (p *PaymentXendit) CreatePaymentRequestEwallet(
	currency payment_request.PaymentRequestCurrency,
	referenceId string,
	amountPrice float64,
	desc payment_request.NullableString,
	paymentMethod payment_request.PaymentMethodParameters,
	customer map[string]interface{},
) (map[string]interface{}, error) {
	paymentRequestParameter := payment_request.PaymentRequestParameters{
		Currency:      currency,
		ReferenceId:   &referenceId,
		PaymentMethod: &paymentMethod,
		Description:   desc,
		Customer:      customer,
		Amount:        &amountPrice,
	}

	res, _, err := p.xenditClient.PaymentRequestApi.CreatePaymentRequest(context.Background()).
		PaymentRequestParameters(paymentRequestParameter).
		Execute()
	if err != nil {
		return nil, err
	}

	if len(res.Actions) == 0 {
		return nil, errors.New("nomor anda salah")
	}

	var urlWeb, urlMobile string

	for _, action := range res.Actions {
		if action.UrlType == "WEB" {
			urlWeb = action.GetUrl()
		}
		if action.UrlType == "MOBILE" {
			urlMobile = action.GetUrl()
		}
	}

	response := map[string]interface{}{
		"Id":                res.Id,
		"Created":           res.Created,
		"ReferenceId":       res.ReferenceId,
		"BusinessId":        res.BusinessId,
		"CustomerId":        *res.CustomerId.Get(),
		"Amount":            fmt.Sprintf("%.0f", *res.Amount),
		"Status":            res.Status,
		"RedirectUrlWeb":    urlWeb,
		"RedirectUrlMobile": urlMobile,
		"Type":              res.PaymentMethod.Type,
		"ChannelCode":       res.PaymentMethod.Ewallet.Get().ChannelCode,
	}

	return response, nil
}

func (p *PaymentXendit) GetPaymentReuqestById(paymentId string) {
	resp, r, err := p.xenditClient.PaymentRequestApi.GetPaymentRequestByID(context.Background(), paymentId).
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.GetPaymentRequestByID``: %v\n", err.Error())

		b, _ := json.Marshal(err.FullError())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	b, _ := json.Marshal(resp)
	fmt.Println("Full response:", string(b))

	fmt.Println(resp)
}
