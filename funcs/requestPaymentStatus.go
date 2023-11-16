package funcs

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func CheckPaymentStatus(accessToken, checkoutRequestID, passKey string) {
	url := safaricomBaseURL + stkPushEndpoint

	payload := map[string]interface{}{
		"BusinessShortCode": shortCode,
		"Password":          lipaNaMpesaOnlinePassword(shortCode, passKey),
		"Timestamp":         timestamp(),
		"CheckoutRequestID": "ws_CO_260520211133524545",
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Printf("could not marshal payload to json with error: %v", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Printf("could not create request with error: %v", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("could not send request with error: %v", err)
		return
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("could not decode response body with error: %v", err)
		return
	}

	log.Println(result)
}
