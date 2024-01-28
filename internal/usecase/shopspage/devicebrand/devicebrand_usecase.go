package usecase

import (
	domain_shopsdevicebrand "content-management/domain/shopspage/devicebrand"
	"content-management/helper"
	"net/http"
	"time"
)

type ShopspageDeviceBrandUsecase struct {
	br domain_shopsdevicebrand.ShopspageDeviceBrandRepository
}

func NewShopspageDeviceBrandUsecase(br domain_shopsdevicebrand.ShopspageDeviceBrandRepository) *ShopspageDeviceBrandUsecase {
	return &ShopspageDeviceBrandUsecase{br}
}

func (p *ShopspageDeviceBrandUsecase) GetShopspageDeviceBrands() (*[]domain_shopsdevicebrand.ShopspageDeviceBrand, error) {
	start := time.Now()

	res, err := p.br.GetShopspageDeviceBrands()
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|DEVICEBRAND|GetSHOPSPAGEDEVICEBRANDS",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	defer helper.CreateLog(&helper.Log{
		Event:        "INTERNAL|USECASE|SHOPSPAGE|DEVICEBRAND|GETSHOPSPAGEDEVICEBRANDS",
		StatusCode:   http.StatusOK,
		ResponseTime: time.Since(start),
		Method:       "GET",
		Message:      "Successfully get data shops page device brands",
	}, "info")

	return res, nil
}
