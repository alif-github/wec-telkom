package domain

import (
	"errors"
	"os"
)

type ShopspageBanner struct {
	Image     string `json:"image"`
	TargetURL string `json:"targetURL"`
}

type ShopspageBannerLabel struct {
	TextButton   string
	TextButtonEN string
	TargetURL    string
}

type ShopspageBannerLabelIdResponse struct {
	CTAButton string            `json:"cta_button"`
	CTALink   string            `json:"cta_link"`
	Data      []ShopspageBanner `json:"data"`
}

type ShopspageBannerLabelEnResponse struct {
	CTAButton string            `json:"cta_button"`
	CTALink   string            `json:"cta_link"`
	Data      []ShopspageBanner `json:"data"`
}

type ShopspageBannerLabelResponse struct {
	LangId ShopspageBannerLabelIdResponse `json:"lang_id"`
	LangEn ShopspageBannerLabelEnResponse `json:"lang_en"`
}

type ShopspageBannerResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewShopspageBanner(Image, TargetURL string) (*ShopspageBanner, error) {
	banner := &ShopspageBanner{
		Image:     Image,
		TargetURL: TargetURL,
	}

	if banner.Image == "" || banner.TargetURL == "" {
		return nil, errors.New("banners primary fields can't be empty")
	}

	return banner, nil
}

func NewShopspageBannerLabel(TextButton, TextButtonEN, TargetURL string) (*ShopspageBannerLabel, error) {
	banner := &ShopspageBannerLabel{
		TextButton:   TextButton,
		TextButtonEN: TextButtonEN,
		TargetURL:    TargetURL,
	}

	if banner.TextButton == "" || banner.TextButtonEN == "" {
		return nil, errors.New("banners label primary fields can't be empty")
	}

	return banner, nil
}

func NewShopspageBannerLabelIdResponse(CTAButton, CTALink string, Data *[]ShopspageBanner) (*ShopspageBannerLabelIdResponse, error) {
	banner := &ShopspageBannerLabelIdResponse{
		CTAButton: CTAButton,
		CTALink:   CTALink,
		Data:      *Data,
	}

	return banner, nil
}

func NewShopspageBannerLabelEnResponse(CTAButton, CTALink string, Data *[]ShopspageBanner) (*ShopspageBannerLabelEnResponse, error) {
	banner := &ShopspageBannerLabelEnResponse{
		CTAButton: CTAButton,
		CTALink:   CTALink,
		Data:      *Data,
	}

	return banner, nil
}

func NewShopspageBannerLabelResponse(LangId *ShopspageBannerLabelIdResponse, LangEn *ShopspageBannerLabelEnResponse) (*ShopspageBannerLabelResponse, error) {
	banner := &ShopspageBannerLabelResponse{
		LangId: *LangId,
		LangEn: *LangEn,
	}

	return banner, nil
}

func GetShopsBanner(p ShopspageBanner) (*ShopspageBanner, error) {
	banner := &ShopspageBanner{
		Image:     os.Getenv("ASSETS_URL") + p.Image,
		TargetURL: p.TargetURL,
	}

	if banner.Image == "" || banner.TargetURL == "" {
		return nil, errors.New("shops page banner primary fields can't be empty")
	}

	return banner, nil
}
