package usecase

import (
	domain_shopsinternetpackage "content-management/domain/shopspage/internetpackage"
	"content-management/helper"
	"net/http"
	"os"
	"strings"
	"time"
)

type ShopspageInternetPackageUsecase struct {
	br domain_shopsinternetpackage.ShopspageInternetPackageRepository
}

func NewShopspageInternetPackageUsecase(br domain_shopsinternetpackage.ShopspageInternetPackageRepository) *ShopspageInternetPackageUsecase {
	return &ShopspageInternetPackageUsecase{br}
}

func (p *ShopspageInternetPackageUsecase) GetShopspageInternetPackage(packageType string) (*domain_shopsinternetpackage.ShopspageInternetPackage, error) {
	start := time.Now()

	internetPackage, err := p.br.GetShopspageInternetPackage(packageType)
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|PACKAGEPOSTPAID|GetSHOPSPAGEPACKAGEPOSTPAID",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	packageLabel, err := p.br.GetShopspageInternetPackageLabel(packageType)
	if err != nil {
		helper.CreateLog(&helper.Log{
			Event:        "INTERNAL|USECASE|SHOPSPAGE|PACKAGEPOSTPAID|GetSHOPSPAGEPACKAGEPOSTPAID",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(start),
			Method:       "GET",
			Message:      err.Error(),
		}, "error")

		return nil, err
	}

	resp := &domain_shopsinternetpackage.ShopspageInternetPackage{
		Title: packageLabel.TItle,
		CTA:   packageLabel.CTA,
	}

	if strings.EqualFold(packageType, os.Getenv("PREPAID_LABEL")) {
		resp.Prepaid = *internetPackage
	} else if strings.EqualFold(packageType, os.Getenv("POSTPAID_LABEL")) {
		resp.Postpaid = *internetPackage
	}

	return resp, nil
}
