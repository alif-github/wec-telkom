package http

import (
	domain_promo "content-management/domain/promo"
	"content-management/helper"
	"content-management/helper/exceptions"
	"encoding/json"
	"net/http"
	"time"
)

type PromoHttpDelivery struct {
	Usecase domain_promo.PromoUseCase
}

func NewPromoHttpDelivery(Usecase domain_promo.PromoUseCase) *PromoHttpDelivery {
	return &PromoHttpDelivery{Usecase: Usecase}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (p *PromoHttpDelivery) GetPromo(w http.ResponseWriter, r *http.Request) {
	event := "INTERNAL|DELIVERY|HTTP|PROMOHTTP|GETPROMO|"
	method := r.Method
	reqURI := r.RequestURI
	targetUrl := r.URL.Query().Get("targetUrl")

	if targetUrl == "" {
		defer helper.CreateLog(&helper.Log{
			Event:        event + "VALIDATIONQUERYSTRINGS",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       method,
			URL:          reqURI,
			Response:     exceptions.ErrPayload.Error(),
		}, "error")

		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid value of query parameter.",
			"status":  http.StatusBadRequest,
		})
		return
	}

	res, err := p.Usecase.GetPromoByTargetUrl(targetUrl)
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        event + "GETPROMOBYTARGETURL",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       method,
			URL:          reqURI,
			Response:     err.Error(),
		}, "error")

		respondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
			"toJson":  true,
		})
		return
	}

	respondWithJSON(w, http.StatusOK, res)
}
