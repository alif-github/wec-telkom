package usecase

import (
	domain_categorypoin "content-management/domain/categorypoin"
	domain_listpoin "content-management/domain/listpoin"
	domain_subscriber "content-management/domain/subscriber"
	"content-management/helper"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ListPoinUsecase struct {
	pr domain_listpoin.ListPoinRepository
	cp domain_categorypoin.CategoryPoinRepository
	sr domain_subscriber.SubscriberService
}

func NewListPoinUsecase(pr domain_listpoin.ListPoinRepository, cp domain_categorypoin.CategoryPoinRepository, sv domain_subscriber.SubscriberService) *ListPoinUsecase {
	return &ListPoinUsecase{pr, cp, sv}
}

func (p *ListPoinUsecase) GetListPoin(categoryName, msisdn string) (*[]domain_listpoin.ListPoin, error) {
	event := "INTERNAL|USECASE|LISTPOIN|GETLISTPOIN"
	start := time.Now()
	var subscriberType string
	var err error
	mockSubscriber, _ := strconv.ParseBool(os.Getenv("MOCKSUBSCRIBER"))

	if !mockSubscriber {
		subscriberType, err = p.sr.GetSubscriberType(msisdn)
		if err != nil {
			return nil, err
		}
	} else {
		subscriberType = os.Getenv("MOCKTYPE")
	}

	categoryId, err := p.cp.GetCategoryIdByName(categoryName)
	if err != nil {
		return nil, err
	}

	res, err := p.pr.GetListPoinByCategoryId(categoryId, subscriberType)
	if err != nil {
		return nil, err
	}

	defer helper.CreateLog(&helper.Log{
		Event:        event,
		StatusCode:   http.StatusOK,
		ResponseTime: time.Since(start),
		Method:       "POST",
		Request:      "category:" + categoryName,
		Message:      "Successfully get data list poin",
	}, "info")

	return res, nil
}

func (p *ListPoinUsecase) DeleteKeywordAutomate() error {
	listCmsPoint, err := p.pr.GetAllListPoin()
	if err != nil {
		return err
	}

	for _, cmsPoint := range *listCmsPoint {
		found, err := p.pr.CheckKeywordInMasterKeyword(cmsPoint.Keyword)
		if err != nil {
			return err
		}

		if !found {
			err := p.pr.DeleteKeyword(cmsPoint.Keyword)
			if err != nil {
				return err
			}
		}
	}

	return err
}
