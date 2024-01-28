package http

import (
	domain_shopsbundlingcard "content-management/domain/shopspage/bundlingcard"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"time"
)

type ShopspageBundlingCardHttpDelivery struct {
	Usecase domain_shopsbundlingcard.ShopspageBundlingCardUsecase
}

type BundlingCardResponse struct {
	domain_shopsbundlingcard.ShopspageBundlingCardResponse
}

func NewShopspageBundlingCardHttpDelivery(Usecase domain_shopsbundlingcard.ShopspageBundlingCardUsecase) *ShopspageBundlingCardHttpDelivery {
	return &ShopspageBundlingCardHttpDelivery{Usecase: Usecase}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (l *ShopspageBundlingCardHttpDelivery) GetShopspageBundlingCard(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	bundlingcards, err := l.Usecase.GetShopspageBundlingCard()
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|HTTP|SHOPSPAGE|BUNDLINGCARD|GETSHOPSPAGEBUNDLINGCARD",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       r.Method,
			URL:          r.RequestURI,
			Response:     err,
		}, "error")

		res := domain_shopsbundlingcard.ShopspageBundlingCardResponse{Status: false, Message: err.Error()}
		respondWithJSON(w, exceptions.MapToHttpStatusCode(err), res)
		return
	}

	res := domain_shopsbundlingcard.ShopspageBundlingCardResponse{
		Status:  true,
		Message: "Successfully get shops page bundling card",
		Data:    bundlingcards,
	}

	respondWithJSON(w, http.StatusOK, res)
}
