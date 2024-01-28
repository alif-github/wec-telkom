package usecase

import (
	domain_shopsbundlingcard "content-management/domain/shopspage/bundlingcard"
	"content-management/helper"
	"net/http"
	"time"
)

type ShopspageBundlingCardUsecase struct {
	br domain_shopsbundlingcard.ShopspageBundlingCardRepository
}

func NewShopspageBundlingCardUsecase(br domain_shopsbundlingcard.ShopspageBundlingCardRepository) *ShopspageBundlingCardUsecase {
	return &ShopspageBundlingCardUsecase{br}
}

func (p *ShopspageBundlingCardUsecase) GetShopspageBundlingCard() (*domain_shopsbundlingcard.ShopspageBundlingCTAResponse, error) {
	start := time.Now()

	bundling, err := p.br.GetShopspageBundlingCard()
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|BUNDLINGCARD|GetSHOPSPAGEBUNDLINGCARD",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	label, err := p.br.GetShopspageBundlingCardLabel()
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|BUNDLINGCARD|GetSHOPSPAGEBUNDLINGCARDLABEL",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	resp := &domain_shopsbundlingcard.ShopspageBundlingCTAResponse{
		CTA:        *label,
		DeviceList: *bundling,
	}

	defer helper.CreateLog(&helper.Log{
		Event:        "INTERNAL|USECASE|SHOPSPAGE|BUNDLINGCARD|GETSHOPSPAGEBUNDLINGCARDS",
		StatusCode:   http.StatusOK,
		ResponseTime: time.Since(start),
		Method:       "GET",
		Message:      "Successfully get data shops page device brands",
	}, "info")

	return resp, nil
}
