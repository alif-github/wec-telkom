package usecase

import (
	domain_promo "content-management/domain/promo"
	"content-management/helper"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type PromoUseCase struct {
	pr domain_promo.PromoRepository
}

func NewPromoUseCase(pr domain_promo.PromoRepository) *PromoUseCase {
	return &PromoUseCase{pr}
}

func (p *PromoUseCase) GetPromoByTargetUrl(targetUrl string) (*domain_promo.PromoResponse, error) {
	var res *domain_promo.Promo
	var err error
	event := "INTERNAL|USECASE|PROMOUSECASE|GETPROMOBYTARGETURL|"
	date := time.Now().Format("2006-01-02 15:04:05")
	isLatest := false

	res, err = p.pr.GetPromoByTargetUrl(targetUrl, date, isLatest)
	if err != nil {
		return nil, err
	}

	if res.Title == "" {
		helper.StringLog("info", event+"GETPROMOBYTARGETURL|LATEST TARGETURL:"+targetUrl+"|DATE:"+date)
		isLatest = true
		res, err = p.pr.GetPromoByTargetUrl(targetUrl, date, isLatest)
		if err != nil {
			return nil, err
		}

		if res.Title == "" {
			return nil, errors.New("error couldn't get data promo")
		}
	}

	mapping := mappingResponse(res)

	response := &domain_promo.PromoResponse{
		Message:   "Successfully get data promo",
		Id:        res.ContentId,
		Category:  res.Category,
		IsEnded:   mapping.IsEnded,
		TypePromo: res.TypePromo,
		ToJson:    false,
		LangEn:    mapping.LangEn,
		LangId:    mapping.LangId,
	}

	defer helper.CreateLog(&helper.Log{
		Event:        event + "SUCCESS",
		StatusCode:   http.StatusOK,
		ResponseTime: time.Since(time.Now()),
		Method:       "GET",
		Request:      "targetUrl: " + targetUrl + " date:" + date,
		Message:      "successfully get data promo",
	}, "info")

	return response, nil
}

func mappingResponse(res *domain_promo.Promo) *domain_promo.MappingResponse {
	var isEnded bool
	var incomingText, incomingTextEn, incomingDate, startDate, endDate string
	dateNow := time.Now().Format("01-02-2006")
	status := helper.GetStatus(res.Status)
	startDate = helper.ChangeDatePosition(*res.StartDate)
	endDate = helper.ChangeDatePosition(*res.EndDate)
	image := fmt.Sprintf("assets/%s", res.Image)

	if res.IncomingDate != "" && res.IncomingText != "" {
		splitDateDash := strings.Split(res.IncomingDate, "-")
		splitDateDash[0], splitDateDash[2] = splitDateDash[2], splitDateDash[0]
		incomingDate = strings.Join(splitDateDash, "-")
		incomingDate += " " + res.StartTime

		incomingText = res.IncomingText
		incomingTextEn = res.IncomingTextEN
	}

	if res.UsingTime == "1" {
		startDate += " " + res.StartTime
		endDate += " " + res.EndTime
		dateNow = time.Now().Format("01-02-2006 15:04:05")
	}

	if endDate != "" {
		if dateNow > endDate {
			isEnded = true
		}
	}

	var tags []string
	json.Unmarshal([]byte(res.Tags), &tags)

	langEn := map[string]interface{}{
		"ML":            res.ML,
		"image":         image,
		"subcategory":   res.Subcategory,
		"page_title":    res.PageTitleEN,
		"title":         res.TitleEN,
		"slug":          res.Slug,
		"description":   res.DescriptionEN,
		"button_text":   res.ButtonTextEN,
		"offering_text": res.OfferingTextEN,
		"start_date":    startDate,
		"end_date":      endDate,
		"terms_title":   res.TermsTitleEN,
		"terms":         res.TermsEN,
		"load_text":     res.LoadTextEN,
		"status":        status,
		"period_text":   res.PeriodTextEN,
		"hide_text":     res.HideTextEN,
		"incoming_text": incomingTextEn,
		"incoming_date": incomingDate,
		"tags_title":    res.TagsTitleEN,
		"tags":          tags,
	}

	langId := map[string]interface{}{
		"ML":            res.ML,
		"image":         image,
		"subcategory":   res.Subcategory,
		"page_title":    res.PageTitle,
		"title":         res.Title,
		"slug":          res.Slug,
		"description":   res.Description,
		"button_text":   res.ButtonText,
		"offering_text": res.OfferingText,
		"start_date":    startDate,
		"end_date":      endDate,
		"terms_title":   res.TermsTitle,
		"terms":         res.Terms,
		"load_text":     res.LoadText,
		"status":        status,
		"period_text":   res.PeriodText,
		"hide_text":     res.HideText,
		"incoming_text": incomingText,
		"incoming_date": incomingDate,
		"tags_title":    res.TagsTitle,
		"tags":          tags,
	}

	resp := &domain_promo.MappingResponse{
		LangEn:  langEn,
		LangId:  langId,
		IsEnded: isEnded,
	}

	return resp
}
