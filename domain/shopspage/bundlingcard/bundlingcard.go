package domain

import (
	"content-management/helper"
	"errors"
	"os"
)

type ShopspageBundlingCard struct {
	Title         string `json:"title"`
	TargetURL     string `json:"target_url"`
	Price         string `json:"real_price"`
	PriceDiscount string `json:"price_after_discount"`
	Description   string `json:"desc"`
	Date          string `json:"date"`
	Image         string `json:"image_url"`
}

type ShopspageBundlingCardCTA struct {
	Description string                    `json:"desc_cta"`
	Subtitle    string                    `json:"sub_title"`
	TargetURL   string                    `json:"url"`
	CTALabel    ShopspageBundlingCTALabel `json:"ctaLabel"`
	CTATitle    ShopspageBundlingCTATitle `json:"title"`
}

type ShopspageBundlingCTALabel struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}
type ShopspageBundlingCTATitle struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}

type ShopspageBundlingCardResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ShopspageBundlingCTAResponse struct {
	CTA        ShopspageBundlingCardCTA `json:"cta"`
	DeviceList []ShopspageBundlingCard  `json:"deviceList"`
}

func NewShopspageBundlingCard(Title, TargetURL, Price, PriceDiscount, Description, Date, Image string) (*ShopspageBundlingCard, error) {
	bundling := &ShopspageBundlingCard{
		Title:         Title,
		TargetURL:     TargetURL,
		Price:         Price,
		PriceDiscount: PriceDiscount,
		Description:   Description,
		Date:          Date,
		Image:         Image,
	}

	if bundling.Title == "" || bundling.TargetURL == "" {
		return nil, errors.New("bundling card primary fields can't be empty")
	}

	return bundling, nil
}

func NewShopspageBundlingCardCTA(Description, Subtitle, TargetURL string, CTALabel ShopspageBundlingCTALabel, CTATitle ShopspageBundlingCTATitle) (*ShopspageBundlingCardCTA, error) {
	bundling := &ShopspageBundlingCardCTA{
		Description: Description,
		Subtitle:    Subtitle,
		TargetURL:   TargetURL,
		CTALabel:    CTALabel,
		CTATitle:    CTATitle,
	}

	if bundling.Description == "" || bundling.TargetURL == "" {
		return nil, errors.New("bundling card cta primary fields can't be empty")
	}

	return bundling, nil
}

func NewShopspageBundlingCTALabel(English, Indonesia string) (*ShopspageBundlingCTALabel, error) {
	label := &ShopspageBundlingCTALabel{
		English:   English,
		Indonesia: Indonesia,
	}

	if label.English == "" || label.Indonesia == "" {
		return nil, errors.New("bundling cta label can't be empty")
	}

	return label, nil
}

func NewShopspageBundlingCTATitle(English, Indonesia string) (*ShopspageBundlingCTATitle, error) {
	label := &ShopspageBundlingCTATitle{
		English:   English,
		Indonesia: Indonesia,
	}

	if label.English == "" || label.Indonesia == "" {
		return nil, errors.New("bundling cta title can't be empty")
	}

	return label, nil
}

func NewShopspageBundlingCTAResponse(CTA ShopspageBundlingCardCTA, DeviceList []ShopspageBundlingCard) (*ShopspageBundlingCTAResponse, error) {
	resp := &ShopspageBundlingCTAResponse{
		CTA:        CTA,
		DeviceList: DeviceList,
	}

	return resp, nil
}

func GetShopsBundlingCard(p ShopspageBundlingCard) (*ShopspageBundlingCard, error) {
	date := helper.DateFormatter(p.Date)
	bundling := &ShopspageBundlingCard{
		Title:         p.Title,
		TargetURL:     p.TargetURL,
		Price:         p.Price,
		PriceDiscount: p.PriceDiscount,
		Description:   p.Description,
		Date:          date,
		Image:         os.Getenv("ASSETS_URL") + p.Image,
	}

	if bundling.Title == "" || bundling.TargetURL == "" {
		return nil, errors.New("shops page bundling primary fields can't be empty")
	}

	return bundling, nil
}
