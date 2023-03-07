package xumm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	BASE_URL = "https://xumm.app/api/v1/platform/%s"
)

type xumm struct {
	client *http.Client
	config XummConfig
}

type XummService interface {
	PingXummServer() (XummPingResponse, int, error)
	XummSignIn() (XummSignInResponse, int, error)
	Payment(request PaymentRequest) (XummSignInResponse, int, error)
}

type XummPingResponse struct {
	Pong bool `json:"pong"`
	Auth struct {
		Application struct {
			UUIDv4     string `json:"uuidv4"`
			Name       string `json:"name"`
			Webhookurl string `json:"webhookurl"`
			Disabled   int    `json:"disabled"`
		} `json:"application"`
		Call struct {
			UUIDv4 string `json:"uuidv4"`
		} `json:"call"`
	} `json:"auth"`
}

type XummmSignInRequest struct {
	UserToken string `json:"user_token"`
	Txjson    struct {
		TransactionType string `json:"TransactionType"`
	} `json:"txjson"`
}

type XummPaymentRequest struct {
	Txjson struct {
		TransactionType string `json:"TransactionType"`
		Destination     string `json:"Destination"`
		DestinationTag  int64  `json:"DestinationTag"`
		Amount          string `json:"Amount"`
	} `json:"txjson"`
}

type PaymentRequest struct {
	Destination    string `json:"address"`
	DestinationTag int64  `json:"destination_tag"`
	Amount         string `json:"amount"`
}

type XummSignInResponse struct {
	UUID string `json:"uuid"`
	Refs struct {
		QRPNG           string `json:"qr_png"`
		WebsocketStatus string `json:"websocket_status"`
	} `json:"refs"`
	Pushed bool `json:"pushed"`
}

type XummConfig struct {
	XUMM_API_KEY    string `json:"xumm_api_key"`
	XUMM_API_SECRET string `json:"xumm_api_secret"`
}

func NewXummService() XummService {
	jsonFile, _ := ioutil.ReadFile("./config/config.json")
	var payload XummConfig
	json.Unmarshal(jsonFile, &payload)
	client := &http.Client{}
	return xumm{
		client: client,
		config: payload,
	}
}

func (x xumm) generateHeaders(r *http.Request) *http.Request {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-API-Key", x.config.XUMM_API_KEY)
	r.Header.Add("X-API-Secret", x.config.XUMM_API_SECRET)
	return r
}
