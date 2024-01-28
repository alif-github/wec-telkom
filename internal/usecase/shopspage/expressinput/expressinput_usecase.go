package usecase

import (
	domain_shopsexpressinput "content-management/domain/shopspage/expressinput"
	"content-management/helper"
	"net/http"
	"time"
)

type ShopspageExpressInputUsecase struct {
	br domain_shopsexpressinput.ShopspageExpressInputRepository
}

func NewShopspageDeviceBrandUsecase(br domain_shopsexpressinput.ShopspageExpressInputRepository) *ShopspageExpressInputUsecase {
	return &ShopspageExpressInputUsecase{br}
}

func (p *ShopspageExpressInputUsecase) GetShopspageExpressInput() (*[]domain_shopsexpressinput.ShopspageExpressInput, error) {
	start := time.Now()

	res, err := p.br.GetShopspageExpressInput()
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|EXPRESSINPUT|GetSHOPSPAGEEXPRESSINPUTS",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	defer helper.CreateLog(&helper.Log{
		Event:        "INTERNAL|USECASE|SHOPSPAGE|EXPRESSINPUT|GETSHOPSPAGEEXPRESSINPUTS",
		StatusCode:   http.StatusOK,
		ResponseTime: time.Since(start),
		Method:       "GET",
		Message:      "Successfully get data shops page express input",
	}, "info")

	return res, nil
}
