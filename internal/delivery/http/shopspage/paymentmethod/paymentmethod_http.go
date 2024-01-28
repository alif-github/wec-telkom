package http

import (
	domain_shopspaymentmethod "content-management/domain/shopspage/paymentmethod"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"time"
)

type ShopspagePaymentMethodHttpDelivery struct {
	Usecase domain_shopspaymentmethod.ShopspagePaymentMethodUsecase
}

func NewShopspagePaymentMethodHttpDelivery(Usecase domain_shopspaymentmethod.ShopspagePaymentMethodUsecase) *ShopspagePaymentMethodHttpDelivery {
	return &ShopspagePaymentMethodHttpDelivery{Usecase: Usecase}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (l *ShopspagePaymentMethodHttpDelivery) GetShopspagePaymentMethods(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	payments, err := l.Usecase.GetShopspagePaymentMethod()
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|HTTP|SHOPSPAGE|PAYMENTMETHOD|GETSHOPSPAGEPAYMENTMETHOD",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       r.Method,
			URL:          r.RequestURI,
			Response:     err,
		}, "error")

		res := domain_shopspaymentmethod.ShopspagePaymentMethodResponse{Status: false, Message: err.Error()}
		respondWithJSON(w, exceptions.MapToHttpStatusCode(err), res)
		return
	}

	res := domain_shopspaymentmethod.ShopspagePaymentMethodResponse{
		Status:  true,
		Message: "Successfully get shops page payment methods",
		Data:    payments,
	}

	respondWithJSON(w, http.StatusOK, res)
}
