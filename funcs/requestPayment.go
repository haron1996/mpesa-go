package funcs

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type paymentRequestResponse struct {
	ResponseCode        string
	CheckoutRequestID   string
	ResponseDescription string
	CustomerMessage     string
}

func newPaymentRequestResponse(responseCode, checkoutRequestID, responseDescription, customerMessage string) *paymentRequestResponse {
	return &paymentRequestResponse{
		ResponseCode:        responseCode,
		CheckoutRequestID:   checkoutRequestID,
		ResponseDescription: responseDescription,
		CustomerMessage:     customerMessage,
	}
}

func RequestPayment(accessToken, passKey, callback string) (*paymentRequestResponse, error) {

	url := safaricomBaseURL + stkPushEndpoint

	payload := map[string]interface{}{
		"BusinessShortCode": shortCode,
		"Password":          lipaNaMpesaOnlinePassword(shortCode, passKey),
		"Timestamp":         timestamp(),
		"TransactionType":   lipaNaMpesaOnline,
		"Amount":            amount,
		"PartyA":            phoneNumber,
		"PartyB":            shortCode,
		"PhoneNumber":       phoneNumber,
		"CallBackURL":       callback,
		"AccountReference":  "123456", // Replace with your unique order ID or reference
		"TransactionDesc":   "Payment for XYZ",
		"ResponseType":      "[Cancelled/Completed]",
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal payload to json with error: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, fmt.Errorf("could not create request with error: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request with error: %v", err)
	}

	defer resp.Body.Close()

	// You can handle the response as needed
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Printf("could not read response body with error: %v", err)
	// 	return
	// }

	//fmt.Println("STK Push Response:", string(body))

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("could not decode response with error: %v", err)
	}

	log.Printf("request payment result: %v", result)

	// check response code
	responseCode, ok := result["ResponseCode"].(string)
	if !ok {
		return nil, errors.New("responseCode not found in response")
	}

	if responseCode != "0" {
		return nil, fmt.Errorf("payment could not be proccessed with code: %s", responseCode)
	}

	// check checkout request id
	checkoutRequestID, ok := result["CheckoutRequestID"].(string)
	if !ok {
		return nil, errors.New("checkout request id not found in response")
	}

	// check response description
	responseDescription, ok := result["ResponseDescription"].(string)
	if !ok {
		return nil, errors.New("response description not found in response")
	}

	// check customer message
	customerMessage, ok := result["CustomerMessage"].(string)
	if !ok {
		return nil, errors.New("customer message not found")
	}

	response := newPaymentRequestResponse(responseCode, checkoutRequestID, responseDescription, customerMessage)

	return response, nil
}

func lipaNaMpesaOnlinePassword(shortCode, passKey string) string {
	time := timestamp()
	data := shortCode + passKey + time
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func timestamp() string {
	return "20" + time.Now().UTC().Format("060102150405")
}
