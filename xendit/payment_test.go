package xendit

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/RaihanMalay21/api-service-riors/xendit/ewallet"
	"github.com/RaihanMalay21/api-service-riors/xendit/virtualAccount"
	"github.com/xendit/xendit-go/v6"
	payment_request "github.com/xendit/xendit-go/v6/payment_request"
)

func TestPaymentRequest(t *testing.T) {
	xenditClient := xendit.NewClient("xnd_development_8jTP7KPfzhIQqRdwebLJYPYO4rru7R904RL8Kxg96nhLBQSwkknvZcvnEdNXk1")

	currency := payment_request.PAYMENTREQUESTCURRENCY_IDR
	paymentMethodType := payment_request.PAYMENTMETHODTYPE_VIRTUAL_ACCOUNT
	virtualAccountChannelCode := payment_request.VIRTUALACCOUNTCHANNELCODE_BCA
	paymentMethodReusability := payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE

	desc := "Pembayaran web online store riors"
	referenceId := "orderId1193"
	amount_price := float64(60000)

	nullableAmount := payment_request.NewNullableFloat64(&amount_price)
	nullableDesc := payment_request.NewNullableString(&desc)
	expPayment := time.Now().Add(5 * time.Minute)

	paymentVirtualAccountParameters := payment_request.VirtualAccountParameters{
		Amount:      *nullableAmount,
		Currency:    &currency,
		ChannelCode: virtualAccountChannelCode,
		ChannelProperties: payment_request.VirtualAccountChannelProperties{
			CustomerName: "raihan malay",
			ExpiresAt:    &expPayment,
		},
	}

	nullableVirtualAccount := payment_request.NewNullableVirtualAccountParameters(&paymentVirtualAccountParameters)

	paymentMethod := payment_request.PaymentMethodParameters{
		Type:           paymentMethodType,
		Reusability:    paymentMethodReusability,
		Description:    *nullableDesc,
		ReferenceId:    &referenceId,
		VirtualAccount: *nullableVirtualAccount,
	}

	var jenis = "Pakai"
	items := []payment_request.PaymentRequestBasketItem{
		{
			ReferenceId: &referenceId,
			Name:        "pen yellow grey",
			Quantity:    10,
			Price:       26000,
			Category:    "Peralatan sekolah",
			Currency:    string(currency),
			Type:        &jenis,
		},
		{
			ReferenceId: &referenceId,
			Name:        "pencil blue grey",
			Quantity:    20,
			Price:       26000,
			Category:    "Peralatan sekolah",
			Currency:    string(currency),
			Type:        &jenis,
		},
	}

	paymentRequestParameters := payment_request.PaymentRequestParameters{
		Currency:      currency,
		ReferenceId:   &referenceId,
		PaymentMethod: &paymentMethod,
		Description:   *nullableDesc,
		Customer: map[string]interface{}{
			"reference_id": referenceId,
			"type":         "INDIVIDUAL",
			"individual_detail": map[string]interface{}{
				"given_names": "RAIHAN",
				"surname":     "GANTENG",
			},
			"email":         "tnybosy@gmail.com",
			"mobile_number": "+6289524474965",
		},
		Items:  items,
		Amount: &amount_price,
	}

	res, r, err := xenditClient.PaymentRequestApi.CreatePaymentRequest(context.Background()).
		PaymentRequestParameters(paymentRequestParameters).
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.CreatePaymentRequest``: %v\n", err.Error())

		b, _ := json.Marshal(err.FullError())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	fmt.Println("Response from `PaymentRequestApi.CreatePaymentRequest`: ", *res.PaymentMethod.VirtualAccount.Get().ChannelProperties.VirtualAccountNumber)
}

// func TestCreateVirtualAccount(t *testing.T) {
// 	loc, err := time.LoadLocation("Asia/Jakarta")
// 	if err != nil {
// 		fmt.Println("Error loading location:", err)
// 		return
// 	}

// 	// Ambil waktu saat ini dalam zona waktu yang diinginkan
// 	currentTime := time.Now().In(loc)
// 	expTime := currentTime.Add(1 * time.Hour)
// 	expirationDate := expTime.Format("2006-01-02T15:04:05-07:00")

// 	requestBody := virtualAccount.VirtualAccount{
// 		ExternalId:     "83023087420222",
// 		BankCode:       "BNI",
// 		Name:           "Raihan Ganteng",
// 		Country:        "ID",
// 		Currency:       "IDR",
// 		IsSingleUse:    true,
// 		IsClosed:       true,
// 		ExpectedAmount: 60000,
// 		ExpirationDate: expirationDate,
// 	}

// 	jsonBody, err := json.Marshal(requestBody)
// 	if err != nil {
// 		fmt.Println("error marshal request body: ", err)
// 		return
// 	}

// 	req, err := http.NewRequest("POST", "https://api.xendit.co/callback_virtual_accounts", bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		fmt.Println("Error create new request: ", err)
// 		return
// 	}

// 	apiKey := os.Getenv("XENDIT-SECRET-KEY")
// 	req.SetBasicAuth(apiKey, "")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("err excute request: ", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Failed to create virtual account. Status Code: %d\n", resp.StatusCode)
// 		body, _ := io.ReadAll(resp.Body)
// 		fmt.Println("Response body:", string(body))
// 		return
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response: ", err)
// 		return
// 	}

// 	var response virtualAccount.ResponseCreateVA
// 	if err := json.Unmarshal(body, &response); err != nil {
// 		fmt.Println("error unmarchall repsonse: ", err)
// 		return
// 	}

// 	fmt.Printf("Virtual account created: %+v\n", response)
// }

func ExprireTimeVA() (string, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}

	currentTime := time.Now().In(loc)

	return currentTime.Add(1 * time.Hour).Format("2006-01-02T15:04:05-07:00"), nil
}

// func TestRequsetVaByModule(t *testing.T) {
// 	var response = make(map[string]interface{})

// 	bank := virtualAccount.ChannelCodePERMATA()

// 	expVA, err := ExprireTimeVA()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	virtualAccount := &VA.VirtualAccount{
// 		ExternalId:     "0312122683712",
// 		BankCode:       bank.BankCode,
// 		Name:           "Programer Baik Hati",
// 		Country:        "ID",
// 		Currency:       "IDR",
// 		IsSingleUse:    true,
// 		IsClosed:       true,
// 		ExpectedAmount: 100000,
// 		ExpirationDate: expVA,
// 	}

// 	resCreateVa, StatusCode := CreateVirtualAccount(virtualAccount, response)

// 	fmt.Println(response)
// 	fmt.Println("response virtual account: ", resCreateVa)
// 	fmt.Println("response Status Code: ", StatusCode)
// }

func TestPayment(t *testing.T) {
	var XC PaymentXendit
	XC.XenditClient()

	exp := time.Now().Add(5 * time.Minute)
	channelProperties := virtualAccount.CreateChannelPropertiesVA("Malay", exp)
	price := float64(1000000)
	amount := payment_request.NewNullableFloat64(&price)
	currency := payment_request.PAYMENTREQUESTCURRENCY_IDR
	channelCode := virtualAccount.ChannelCodePERMATA()

	resultVirtualAccout := virtualAccount.CreateVirtualAccount(*amount, currency, channelCode, channelProperties)

	method := payment_request.PAYMENTMETHODTYPE_VIRTUAL_ACCOUNT
	reusability := payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE
	desc := "bayar utang"
	deskipsi := payment_request.NewNullableString(&desc)
	referensiId := "Id387867499349708033"

	paymentMethod := virtualAccount.CreatePaymentMethodVA(method, reusability, *deskipsi, referensiId, resultVirtualAccout)

	// products := []domain.Product{
	// 	{
	// 		Id:          902901282,
	// 		ProductName: "shirt long hand",
	// 		Price:       50000,
	// 	},
	// 	{
	// 		Id:          902989322,
	// 		ProductName: "shirt short hand",
	// 		Price:       90000,
	// 	},
	// }

	// var items []payment_request.PaymentRequestBasketItem
	// for _, product := range products {
	// 	price := 5 * product.Price
	// 	item := CreateBasketItem(fmt.Sprintf("%d", product.Id), product.ProductName, float64(5), float64(price), "appreal", "IDR", "used")
	// 	items = append(items, item)
	// }

	customerId := "riors73424546"
	nullableId := payment_request.NewNullableString(&customerId)
	// customer := map[string]interface{}{
	// 	"reference_id": "hfuheowhf893793",
	// 	"type":         "INDIVIDUAL",
	// 	"individual_detail": map[string]interface{}{
	// 		"given_names": "RAIHAN",
	// 		"surname":     "GANTENG",
	// 	},
	// 	"email": "tnybosy@gmail.com",
	// }

	result, err := XC.CreatePaymentRequestVA(currency, referensiId, price, *deskipsi, paymentMethod, *nullableId)
	if err != nil {
		b, _ := json.Marshal(err.Error())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))
		return
	}

	fmt.Println(result)
}

func TestPaymentEwallet(t *testing.T) {
	var XC PaymentXendit
	XC.XenditClient()

	successUrl := "http://localhost:8080"
	failureUrl := "http://localhost:8080"
	cancelUrl := "http://localhost:8080"
	mobileNumber := "+6289524474969"
	// exp := time.Now().Add(5 * time.Minute)
	channelProperties := ewallet.CreateChannelPropertiesEwallet(&successUrl, &failureUrl, &cancelUrl, &mobileNumber)
	price := float64(900000)
	// amount := payment_request.NewNullableFloat64(&price)
	currency := payment_request.PAYMENTREQUESTCURRENCY_IDR
	channelCode := ewallet.ChannelCodeLINKAJA()

	resultEwallet := ewallet.CreateEwallet(channelCode, channelProperties)

	method := payment_request.PAYMENTMETHODTYPE_EWALLET
	reusability := payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE
	desc := "bayar utang"
	referensiId := "Id387722897236"

	paymentMethod := ewallet.CreatePaymentMethodEWALLET(method, reusability, *payment_request.NewNullableString(&desc), referensiId, resultEwallet)

	// products := []domain.Product{
	// 	{
	// 		Id:          902901282,
	// 		ProductName: "shirt long hand",
	// 		Price:       50000,
	// 	},
	// 	{
	// 		Id:          902989322,
	// 		ProductName: "shirt short hand",
	// 		Price:       90000,
	// 	},
	// }

	// var items []payment_request.PaymentRequestBasketItem
	// for _, product := range products {
	// 	price := 5 * product.Price
	// 	item := CreateBasketItem(fmt.Sprintf("%d", product.Id), product.ProductName, float64(5), float64(price), "appreal", "IDR", "used")
	// 	items = append(items, item)
	// }

	// customerId := "riors73425"
	customer := map[string]interface{}{
		"reference_id": referensiId,
		"type":         "INDIVIDUAL",
		"individual_detail": map[string]interface{}{
			"given_names": "RAIHAN",
			"surname":     "GANTENG",
		},
		"email":         "tnybosy@gmail.com",
		"mobile_number": "+6289524474965",
	}

	fmt.Println(currency)
	fmt.Println(referensiId)
	fmt.Println(*payment_request.NewNullableString(&desc))
	// fmt.Println(items)
	fmt.Println(customer)

	fmt.Println(XC)
	result, err := XC.CreatePaymentRequestEwallet(
		currency,
		referensiId,
		price,
		*payment_request.NewNullableString(&desc),
		paymentMethod,
		// items,
		customer,
	)
	if err != nil {
		b, _ := json.Marshal(err.Error())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))
		return
	}

	fmt.Println(result)
}

func TestGetPaymentById(t *testing.T) {
	var XC PaymentXendit
	XC.XenditClient()
	id := "pr-4405bac1-0048-4405-91fc-42568c573c03"
	XC.GetPaymentReuqestById(id)
}
