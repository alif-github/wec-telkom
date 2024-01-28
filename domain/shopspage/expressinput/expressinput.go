package domain

import (
	"content-management/helper"
	"errors"
	"os"
)

type ShopspageExpressInput struct {
	Image       string                     `json:"banner"`
	Description string                     `json:"desc"`
	Date        string                     `json:"date"`
	TargetURL   string                     `json:"url"`
	Title       ShopspageExpressInputTitle `json:"title"`
}

type ShopspageExpressInputTitle struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}

type ShopspageExpressInputResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewShopspageExpressInput(Image, Description, Date string, TargetURL string, Title ShopspageExpressInputTitle) (*ShopspageExpressInput, error) {
	input := &ShopspageExpressInput{
		Image:       Image,
		Description: Description,
		Date:        Date,
		TargetURL:   TargetURL,
		Title:       Title,
	}

	if input.Image == "" || input.TargetURL == "" {
		return nil, errors.New("express input primary fields can't be empty")
	}

	return input, nil
}

func GetShopsExpressInput(p ShopspageExpressInput) (*ShopspageExpressInput, error) {
	date := helper.DateFormatter(p.Date)
	input := &ShopspageExpressInput{
		Image:       os.Getenv("ASSETS_URL") + p.Image,
		Description: p.Description,
		Date:        date,
		TargetURL:   p.TargetURL,
		Title:       p.Title,
	}

	if input.TargetURL == "" {
		return nil, errors.New("shops page express input primary fields can't be empty")
	}

	return input, nil
}

func GetShopsExpressInputTitle(p ShopspageExpressInputTitle) (*ShopspageExpressInputTitle, error) {
	title := &ShopspageExpressInputTitle{
		Indonesia: p.Indonesia,
		English:   p.English,
	}

	if title.Indonesia == "" || title.English == "" {
		return nil, errors.New("shops page express input title primary fields can't be empty")
	}

	return title, nil
}
