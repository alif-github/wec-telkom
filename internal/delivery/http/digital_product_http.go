package http

import (
	domain "content-management/domain/digital-product"
	"content-management/helper"
	"content-management/helper/exceptions"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type DigitalProductHttpDelivery struct {
	Usecase domain.DigitalProductUseCase
}

func NewDigitalProductDelivery(Usecase domain.DigitalProductUseCase) *DigitalProductHttpDelivery {
	return &DigitalProductHttpDelivery{Usecase: Usecase}
}

func (dp *DigitalProductHttpDelivery) GetCredit(w http.ResponseWriter, r *http.Request) {

	creditOffers, err := dp.Usecase.GetCredit()
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "digital-product/credit",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "GET",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")

		respondWithJSON(w, http.StatusBadGateway, creditOffers)
	}

	respondWithJSON(w, http.StatusOK, creditOffers)
}
func (dp *DigitalProductHttpDelivery) FindCredit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, validID := vars["creditid"]
	if !validID {
		defer helper.CreateLog(&helper.Log{
			Event:        "digital-product/credit/reduce-stock",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "PATCH",
			URL:          "nil",
			Response:     "id invalid",
		}, "error")

		respondWithJSON(w, http.StatusBadRequest, "id invalid")
		return
	}
	creditOffers, err := dp.Usecase.FindCredit(id)
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "digital-product/credit/{id}",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "GET",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")

		respondWithJSON(w, http.StatusBadGateway, creditOffers)
	}

	respondWithJSON(w, http.StatusOK, creditOffers)
}

func (dp *DigitalProductHttpDelivery) ReduceStock(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, validID := vars["id"]
	if !validID {
		defer helper.CreateLog(&helper.Log{
			Event:        "digital-product/credit/reduce-stock",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "PATCH",
			URL:          "nil",
			Response:     "id invalid",
		}, "error")

		respondWithJSON(w, http.StatusBadRequest, "id invalid")
		return
	}
	err := dp.Usecase.ReduceStock(id)
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "digital-product/credit/reduce-stock",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "PATCH",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, "success")
}
