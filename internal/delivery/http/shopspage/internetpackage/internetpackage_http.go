package http

import (
	domain_shopsinternetpackage "content-management/domain/shopspage/internetpackage"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type ShopspageInternetPackageHttpDelivery struct {
	Usecase domain_shopsinternetpackage.ShopspageInternetPackageUsecase
}

type InternetPackageResponse struct {
	domain_shopsinternetpackage.ShopspageInternetPackageResponse
}

type InternetPackageRequest struct {
	domain_shopsinternetpackage.ShopspageInternetPackageRequest
}

func NewShopspageInternetPackageHttpDelivery(Usecase domain_shopsinternetpackage.ShopspageInternetPackageUsecase) *ShopspageInternetPackageHttpDelivery {
	return &ShopspageInternetPackageHttpDelivery{Usecase: Usecase}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (l *ShopspageInternetPackageHttpDelivery) GetShopspagePrepaid(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	prepaid, err := l.Usecase.GetShopspageInternetPackage(os.Getenv("PREPAID_LABEL"))
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|HTTP|SHOPSPAGE|INTERNETPACKAGE|GETSHOPSPAGEPREPAID",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       r.Method,
			URL:          r.RequestURI,
			Response:     err,
		}, "error")

		res := domain_shopsinternetpackage.ShopspageInternetPackageResponse{Status: false, Message: err.Error()}
		respondWithJSON(w, exceptions.MapToHttpStatusCode(err), res)
		return
	}
	res := domain_shopsinternetpackage.ShopspageInternetPackageResponse{
		Status:  true,
		Message: "Successfully get shops page prepaid internet package",
		Data:    prepaid,
	}

	respondWithJSON(w, http.StatusOK, res)
}

func (l *ShopspageInternetPackageHttpDelivery) GetShopspagePostpaid(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	postpaid, err := l.Usecase.GetShopspageInternetPackage(os.Getenv("POSTPAID_LABEL"))
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|HTTP|SHOPSPAGE|INTERNETPACKAGE|GETSHOPSPAGEPOSTPAID",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       r.Method,
			URL:          r.RequestURI,
			Response:     err,
		}, "error")

		res := domain_shopsinternetpackage.ShopspageInternetPackageResponse{Status: false, Message: err.Error()}
		respondWithJSON(w, exceptions.MapToHttpStatusCode(err), res)
		return
	}
	res := domain_shopsinternetpackage.ShopspageInternetPackageResponse{
		Status:  true,
		Message: "Successfully get shops page postpaid internet package",
		Data:    postpaid,
	}

	respondWithJSON(w, http.StatusOK, res)
}
