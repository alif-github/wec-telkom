package usecase

import (
	domain_shopsbanner "content-management/domain/shopspage/banner"
	"content-management/helper"
	"net/http"
	"time"
)

type ShopspageBannerUsecase struct {
	br domain_shopsbanner.ShopspageBannerRepository
}

func NewShopspageBannerUsecase(br domain_shopsbanner.ShopspageBannerRepository) *ShopspageBannerUsecase {
	return &ShopspageBannerUsecase{br}
}

func (p *ShopspageBannerUsecase) GetShopspageBanners() (*domain_shopsbanner.ShopspageBannerLabelResponse, error) {
	start := time.Now()

	banners, err := p.br.GetShopspageBanners()
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|BANNER|GetSHOPSPAGEBANNERS",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	label, err := p.br.GetShopspageBannerLabel()
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|BANNER|GETSHOPSPAGEBANNERLABEL",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	bannerId, _ := domain_shopsbanner.NewShopspageBannerLabelIdResponse(label.TextButton, label.TargetURL, banners)
	bannerEn, _ := domain_shopsbanner.NewShopspageBannerLabelEnResponse(label.TextButtonEN, label.TargetURL, banners)
	res, _ := domain_shopsbanner.NewShopspageBannerLabelResponse(bannerId, bannerEn)

	defer helper.CreateLog(&helper.Log{
		Event:        "INTERNAL|USECASE|LISTPOIN|GETLISTPOIN",
		StatusCode:   http.StatusOK,
		ResponseTime: time.Since(start),
		Method:       "GET",
		Message:      "Successfully get data shops page banner",
	}, "info")

	return res, nil
}
