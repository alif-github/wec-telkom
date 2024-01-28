package http

import (
	domain_shopsdevicebrand "content-management/domain/shopspage/devicebrand"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"time"
)

type ShopspageDeviceBrandHttpDelivery struct {
	Usecase domain_shopsdevicebrand.ShopspageDeviceBrandUsecase
}

type ListPoinResponse struct {
	domain_shopsdevicebrand.ShopspageDeviceBrandResponse
}

func NewShopspageDeviceBrandHttpDelivery(Usecase domain_shopsdevicebrand.ShopspageDeviceBrandUsecase) *ShopspageDeviceBrandHttpDelivery {
	return &ShopspageDeviceBrandHttpDelivery{Usecase: Usecase}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (l *ShopspageDeviceBrandHttpDelivery) GetShopspageDeviceBrands(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	devicebrands, err := l.Usecase.GetShopspageDeviceBrands()
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|HTTP|SHOPSPAGE|BANNER|GETSHOPSPAGEBANNER",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       r.Method,
			URL:          r.RequestURI,
			Response:     err,
		}, "error")

		res := domain_shopsdevicebrand.ShopspageDeviceBrandResponse{Status: false, Message: err.Error()}
		respondWithJSON(w, exceptions.MapToHttpStatusCode(err), res)
		return
	}

	res := domain_shopsdevicebrand.ShopspageDeviceBrandResponse{
		Status:  true,
		Message: "Successfully get shops page device brands",
		Data:    devicebrands,
	}

	respondWithJSON(w, http.StatusOK, res)
}
