package http

import (
	domain_shopsbanner "content-management/domain/shopspage/banner"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"time"
)

type ShopspageBannerHttpDelivery struct {
	Usecase domain_shopsbanner.ShopspageBannerUsecase
}

type ListPoinResponse struct {
	domain_shopsbanner.ShopspageBannerResponse
}

func NewShopspageBannerHttpDelivery(Usecase domain_shopsbanner.ShopspageBannerUsecase) *ShopspageBannerHttpDelivery {
	return &ShopspageBannerHttpDelivery{Usecase: Usecase}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (l *ShopspageBannerHttpDelivery) GetShopspageBanners(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	banners, err := l.Usecase.GetShopspageBanners()
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|HTTP|SHOPSPAGE|BANNER|GETSHOPSPAGEBANNER",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       r.Method,
			URL:          r.RequestURI,
			Response:     err,
		}, "error")

		res := domain_shopsbanner.ShopspageBannerResponse{Status: false, Message: err.Error()}
		respondWithJSON(w, exceptions.MapToHttpStatusCode(err), res)
		return
	}

	res := domain_shopsbanner.ShopspageBannerResponse{
		Status:  true,
		Message: "Successfully get shops page banners",
		Data:    banners,
	}

	respondWithJSON(w, http.StatusOK, res)
}
