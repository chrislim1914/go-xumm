package api

import (
	"encoding/json"
	"github/chrislim1914/go-xumm/xumm"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func jsonResponseHandler(w http.ResponseWriter, data interface{}, statuscode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(data)
}

func jsonErrorResponseHandler(w http.ResponseWriter, errMsg string, statuscode int) {
	newErr := make(map[string]string)
	newErr["error"] = errMsg
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(newErr)
}

func CheckXummServer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	newxsrv := xumm.NewXummService()
	res, status, err := newxsrv.PingXummServer()
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), status)
		return
	}
	jsonResponseHandler(w, res, status)
}

func SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	newxsrv := xumm.NewXummService()
	res, status, err := newxsrv.XummSignIn()
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), status)
		return
	}
	jsonResponseHandler(w, res, status)
}

func SendPayment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request xumm.PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	newxsrv := xumm.NewXummService()
	res, status, err := newxsrv.Payment(request)
	if err != nil {
		jsonErrorResponseHandler(w, err.Error(), status)
		return
	}
	jsonResponseHandler(w, res, status)
}
