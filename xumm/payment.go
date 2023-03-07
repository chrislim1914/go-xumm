package xumm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (x xumm) Payment(request PaymentRequest) (XummSignInResponse, int, error) {
	var jsonResponse XummSignInResponse
	url := fmt.Sprintf(BASE_URL, "payload")

	payload := XummPaymentRequest{}
	payload.Txjson.TransactionType = "Payment"
	payload.Txjson.Destination = request.Destination
	payload.Txjson.DestinationTag = request.DestinationTag
	payload.Txjson.Amount = request.Amount

	byteRequest, _ := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(byteRequest))
	if err != nil {
		return jsonResponse, http.StatusInternalServerError, err
	}
	req.Header = x.generateHeaders(req).Header
	r, err := x.client.Do(req)
	if err != nil {
		return jsonResponse, http.StatusInternalServerError, err
	}
	defer r.Body.Close()
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return jsonResponse, http.StatusInternalServerError, err
	}
	if r.StatusCode != 200 {
		fmt.Println(string(responseBody))
		newerr := errors.New("something went wrong")
		return jsonResponse, r.StatusCode, newerr
	}
	err = json.Unmarshal(responseBody, &jsonResponse)
	if err != nil {
		return jsonResponse, http.StatusInternalServerError, err
	}
	return jsonResponse, r.StatusCode, nil
}
