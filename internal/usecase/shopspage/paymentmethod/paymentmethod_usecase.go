package usecase

import (
	domain_shopspaymentmethod "content-management/domain/shopspage/paymentmethod"
	"content-management/helper"
	"net/http"
	"time"
)

type ShopspagePaymentMethodUsecase struct {
	br domain_shopspaymentmethod.ShopspagePaymentMethodRepository
}

func NewShopspagePaymentMethodUsecase(br domain_shopspaymentmethod.ShopspagePaymentMethodRepository) *ShopspagePaymentMethodUsecase {
	return &ShopspagePaymentMethodUsecase{br}
}

func (p *ShopspagePaymentMethodUsecase) GetShopspagePaymentMethod() (*[]domain_shopspaymentmethod.ShopspagePaymentMethod, error) {
	start := time.Now()

	res, err := p.br.GetShopspagePaymentMethod()
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|PAYMENTMETHOD|GetSHOPSPAGEPAYMENTMETHODS",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	defer helper.CreateLog(&helper.Log{
		Event:        "INTERNAL|USECASE|SHOPSPAGE|PAYMENTMETHOD|GETSHOPSPAGEPAYMENTMETHODS",
		StatusCode:   http.StatusOK,
		ResponseTime: time.Since(start),
		Method:       "GET",
		Message:      "Successfully get data shops page payment method",
	}, "info")

	return res, nil
}
