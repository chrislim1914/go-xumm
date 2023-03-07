package xumm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (x xumm) PingXummServer() (XummPingResponse, int, error) {
	var jsonResponse XummPingResponse
	url := fmt.Sprintf(BASE_URL, "ping")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return jsonResponse, http.StatusInternalServerError, err
	}
	req.Header = x.generateHeaders(req).Header
	r, err := x.client.Do(req)
	if err != nil {
		return jsonResponse, http.StatusInternalServerError, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		newerr := errors.New("something went wrong")
		return jsonResponse, r.StatusCode, newerr
	}
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return jsonResponse, http.StatusInternalServerError, err
	}
	err = json.Unmarshal(responseBody, &jsonResponse)
	if err != nil {
		return jsonResponse, http.StatusInternalServerError, err
	}
	return jsonResponse, r.StatusCode, nil
}
