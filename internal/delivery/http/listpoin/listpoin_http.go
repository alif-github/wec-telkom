package http

import (
	domain_listpoin "content-management/domain/listpoin"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"time"
)

type ListPoinHttpDelivery struct {
	Usecase domain_listpoin.ListPoinUseCase
}

type ListPoinRequest struct {
	domain_listpoin.ListPoinRequest
}

type ListPoinResponse struct {
	domain_listpoin.ListPoinResponse
}

func NewListPoinHttpDelivery(Usecase domain_listpoin.ListPoinUseCase) *ListPoinHttpDelivery {
	return &ListPoinHttpDelivery{Usecase: Usecase}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (l *ListPoinHttpDelivery) GetListPoin(w http.ResponseWriter, r *http.Request) {
	event := "INTERNAL|DELIVERY|HTTP|LISTPOIN|GETLISTPOIN|"
	start := time.Now()
	var p ListPoinRequest
	method := r.Method
	reqURI := r.RequestURI

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		helper.HttpLog("error", event+"DECODE", method, http.StatusBadRequest, "Invalid Request Payload", time.Since(start), reqURI, "")
		res := domain_listpoin.ListPoinResponse{Status: false, Message: "Invalid Request Payload"}
		respondWithJSON(w, http.StatusBadRequest, res)
	}
	defer r.Body.Close()

	if p.Category == "" {
		helper.HttpLog("error", event+"VALIDATIONPAYLOAD", method, http.StatusBadRequest, "Payload Category Cannot Be Empty", time.Since(start), reqURI, "")
		res := domain_listpoin.ListPoinResponse{Status: false, Message: "Invalid Request Payload"}
		respondWithJSON(w, http.StatusBadRequest, res)
	} else if p.Msisdn == "" {
		helper.HttpLog("error", event+"VALIDATIONPAYLOAD", method, http.StatusBadRequest, "Payload Msisdn Cannot Be Empty", time.Since(start), reqURI, "")
		res := domain_listpoin.ListPoinResponse{Status: false, Message: "Invalid Request Payload"}
		respondWithJSON(w, http.StatusBadRequest, res)
	}

	poin, err := l.Usecase.GetListPoin(p.Category, p.Msisdn)
	if err != nil {
		helper.HttpLog("error", event, method, http.StatusBadRequest, "Invalid Request Payload", time.Since(start), reqURI, "")
		res := domain_listpoin.ListPoinResponse{Status: false, Message: err.Error()}
		respondWithJSON(w, exceptions.MapToHttpStatusCode(err), res)
		return
	}

	res := domain_listpoin.ListPoinResponse{
		Status:  true,
		Message: "Successfully get list poin",
		Data:    poin,
	}

	helper.HttpLog("info", event, method, http.StatusOK, "Successfully Get Data List Poin", time.Since(start), reqURI, "")
	respondWithJSON(w, http.StatusOK, res)
}
