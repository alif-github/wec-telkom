package http

import (
	domain_shopsexpressinput "content-management/domain/shopspage/expressinput"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"time"
)

type ShopspageExpressInputHttpDelivery struct {
	Usecase domain_shopsexpressinput.ShopspageExpressInputUsecase
}

func NewShopspageExpressInputHttpDelivery(Usecase domain_shopsexpressinput.ShopspageExpressInputUsecase) *ShopspageExpressInputHttpDelivery {
	return &ShopspageExpressInputHttpDelivery{Usecase: Usecase}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (l *ShopspageExpressInputHttpDelivery) GetShopspageExpressInput(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	expressinput, err := l.Usecase.GetShopspageExpressInput()
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|HTTP|SHOPSPAGE|EXPRESSINPUT|GETSHOPSPAGEEXPRESSINPUT",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       r.Method,
			URL:          r.RequestURI,
			Response:     err,
		}, "error")

		res := domain_shopsexpressinput.ShopspageExpressInputResponse{Status: false, Message: err.Error()}
		respondWithJSON(w, exceptions.MapToHttpStatusCode(err), res)
		return
	}

	res := domain_shopsexpressinput.ShopspageExpressInputResponse{
		Status:  true,
		Message: "Successfully get shops page express input",
		Data:    expressinput,
	}

	respondWithJSON(w, http.StatusOK, res)
}
