package domain

import (
	"errors"
)

type Promo struct {
	ButtonText     string
	ButtonTextEN   string
	Category       string
	ContentId      string
	Description    string
	DescriptionEN  string
	EndDate        *string
	EndTime        string
	HideText       string
	HideTextEN     string
	Image          string
	IncomingDate   string
	IncomingText   string
	IncomingTextEN string
	LoadText       string
	LoadTextEN     string
	ML             string
	OfferingText   string
	OfferingTextEN string
	PageTitle      string
	PageTitleEN    string
	PeriodText     *string
	PeriodTextEN   *string
	Slug           string
	StartDate      *string
	StartTime      string
	Status         string
	Subcategory    string
	Tags           string
	TagsTitle      string
	TagsTitleEN    string
	TargetURL      string
	Title          string
	TitleEN        string
	TermsTitle     string
	TermsTitleEN   string
	Terms          string
	TermsEN        string
	TypePromo      string
	UsingTime      string
}

type PromoEN struct {
}

type PromoResponse struct {
	Message   string      `json:"message"`
	Category  string      `json:"category,omitempty"`
	Id        string      `json:"id"`
	IsEnded   bool        `json:"isEnded"`
	TypePromo string      `json:"typePromo,omitempty"`
	LangEn    interface{} `json:"lang_en"`
	LangId    interface{} `json:"lang_id"`
	ToJson    bool        `json:"toJson"`
}

type PromotNotFoundResponse struct {
	Message string `json:"message"`
	ToJson  bool   `json:"toJson"`
}

type MappingResponse struct {
	LangEn  interface{}
	LangId  interface{}
	IsEnded bool
}

func NewPromo(ButtonText, ButtonTextEN, Category, Description, DescriptionEN string, EndDate *string, EndTime, HideText, HideTextEN, Image, LoadText, LoadTextEN, ML, OfferingText,
	OfferingTextEN, PageTitle string, PeriodText, PeriodTextEN *string, Slug string, StartDate *string, StarTtime, Status, Title, TermsTitle, TermsTitleEN, Terms, TermsEN, TypePromo, UsingTime string) (*Promo, error) {

	Promo := &Promo{
		ButtonText:     ButtonText,
		ButtonTextEN:   ButtonTextEN,
		Category:       Category,
		Description:    Description,
		DescriptionEN:  DescriptionEN,
		EndDate:        EndDate,
		EndTime:        EndTime,
		HideText:       HideText,
		HideTextEN:     HideTextEN,
		Image:          Image,
		LoadText:       LoadText,
		LoadTextEN:     LoadTextEN,
		ML:             ML,
		OfferingText:   OfferingText,
		OfferingTextEN: OfferingTextEN,
		PageTitle:      PageTitle,
		PeriodText:     PeriodText,
		PeriodTextEN:   PeriodTextEN,
		Slug:           Slug,
		StartDate:      StartDate,
		StartTime:      StarTtime,
		Status:         Status,
		Title:          Title,
		TermsTitle:     TermsTitle,
		TermsTitleEN:   TermsTitleEN,
		Terms:          Terms,
		TermsEN:        TermsEN,
		TypePromo:      TypePromo,
		UsingTime:      UsingTime,
	}

	if Promo.Title == "" || Promo.ButtonText == "" || Promo.OfferingText == "" || Promo.LoadText == "" || Promo.HideText == "" {
		return nil, errors.New("primary fields can't be empty")
	}

	return Promo, nil
}
