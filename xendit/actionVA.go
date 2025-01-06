package xendit

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"net/http"

// 	VA "github.com/RaihanMalay21/api-service-riors/xendit/virtualAccount"
// 	payment_request "github.com/xendit/xendit-go/v6/payment_request"
// )

// func TCreateVirtualAccount(va *VA.VirtualAccount, response map[string]interface{}) (*VA.ResponseCreateVA, int) {
// 	jsonBody, err := json.Marshal(va)
// 	if err != nil {
// 		response["error"] = err.Error()
// 		return nil, http.StatusInternalServerError
// 	}

// 	req, err := http.NewRequest("POST", "https://api.xendit.co/callback_virtual_accounts", bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		response["error"] = err.Error()
// 		return nil, http.StatusInternalServerError
// 	}

// 	apiKey := "xnd_development_8jTP7KPfzhIQqRdwebLJYPYO4rru7R904RL8Kxg96nhLBQSwkknvZcvnEdNXk1"
// 	req.SetBasicAuth(apiKey, "")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		response["error"] = err.Error()
// 		return nil, http.StatusInternalServerError
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		response["error"] = err.Error()
// 		return nil, http.StatusInternalServerError
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		response["error"] = string(body)
// 		return nil, http.StatusBadRequest
// 	}

// 	var resCreateVa VA.ResponseCreateVA
// 	if err := json.Unmarshal(body, &resCreateVa); err != nil {
// 		response["error"] = err.Error()
// 		return nil, http.StatusInternalServerError
// 	}

// 	return &resCreateVa, http.StatusOK
// }

// func CreateEwallet(channelCode payment_request.EWalletChannelCode, channelProperties payment_request.EWalletChannelProperties) {

// }
