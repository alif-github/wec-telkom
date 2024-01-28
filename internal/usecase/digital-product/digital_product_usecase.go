package usecase

import (
	domain "content-management/domain/digital-product"
	"content-management/helper"
	"content-management/helper/exceptions"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type DigitalProductUseCase struct {
	pr domain.DigitalProductRepository
}

func NewDigitalProductUseCase(pr domain.DigitalProductUseCase) *DigitalProductUseCase {
	return &DigitalProductUseCase{pr}
}

func (p *DigitalProductUseCase) GetCredit() (*domain.CreditOffersV2, error) {
	creditResult, err := p.pr.GetCredit()
	if err != nil {
		logrus.Errorf("[digitalProducUsecase.GetCredit] Error on GetCredit: %s", err)
	}

	return creditResult, nil
}

func (p *DigitalProductUseCase) FindCredit(id string) (*domain.ResponseFindCredit, error) {

	creditResult, err := p.pr.FindCredit(id)
	if err != nil {
		logrus.Errorf("[digitalProducUsecase.GetCredit] Error on FindCredit: %s", err)
	}

	return creditResult, nil
}

func (p *DigitalProductUseCase) ReduceStock(packageID string) error {

	err := p.pr.ReduceStock(packageID)
	if err != nil {
		defer helper.CreateLog(&helper.Log{
			Event:        "digital-product/credit/reduce-stock",
			StatusCode:   http.StatusBadRequest,
			ResponseTime: time.Since(time.Now()),
			Method:       "PATCH",
			URL:          "nil",
			Response:     exceptions.ErrPayload.Error(),
		}, "error")
		log.Printf("usecase reduce stock error: %s", err.Error())
	}
	return err
}
