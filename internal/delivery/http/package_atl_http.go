package http

import (
	domain "content-management/domain/atl"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ATLPackageHttpDelivery struct {
	Usecase domain.ATLPackageUseCase
}

func NewATLPackageHttpDelivery(Usecase domain.ATLPackageUseCase) *ATLPackageHttpDelivery {
	return &ATLPackageHttpDelivery{Usecase: Usecase}
}

func (p *ATLPackageHttpDelivery) GetPackage(w http.ResponseWriter, r *http.Request) {
	var payload domain.ATLRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {

		defer helper.CreateLog(&helper.Log{
			Event:        "get-atl-package",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "GET",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")

		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}
	getAtlOffers, err := p.Usecase.GetPackage(payload)

	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "get-atl-package",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "GET",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")
	}

	respondWithJSON(w, http.StatusOK, getAtlOffers)
}

func (p *ATLPackageHttpDelivery) FindPackage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, validID := vars["id"]
	if !validID {
		defer helper.CreateLog(&helper.Log{
			Event:        "get-detail-atl-package",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "PATCH",
			URL:          "nil",
			Response:     "id invalid",
		}, "error")

		respondWithJSON(w, http.StatusBadRequest, "id invalid")
		return
	}

	getAtlOffers, err := p.Usecase.FindPackage(id)
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "get-detail-atl-package",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "GET",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")
	}
	if getAtlOffers.ID == "" {
		defer helper.CreateLog(&helper.Log{
			Event:        "get-detail-atl-package",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "GET",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")
		respondWithJSON(w, http.StatusBadRequest, "id not found")
		return
	}

	respondWithJSON(w, http.StatusOK, getAtlOffers)
}

func (p *ATLPackageHttpDelivery) GetRegion(w http.ResponseWriter, r *http.Request) {

	getAtlOffers, err := p.Usecase.GetRegion()
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "get-region",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "GET",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")
	}
	respondWithJSON(w, http.StatusOK, getAtlOffers)
}
