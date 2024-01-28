package domain

import (
	"content-management/helper"
	"errors"
	"os"
)

type ShopspageDeviceBrand struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"desc"`
	Image       string `json:"banner"`
}

type ShopspageDeviceBrandResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewShopspageDeviceBrand(title, date string, description, image string) (*ShopspageDeviceBrand, error) {
	device := &ShopspageDeviceBrand{
		Title:       title,
		Date:        date,
		Description: description,
		Image:       image,
	}

	if device.Title == "" || device.Date == "" {
		return nil, errors.New("device brand primary fields can't be empty")
	}

	return device, nil
}

func GetShopsBanner(p ShopspageDeviceBrand) (*ShopspageDeviceBrand, error) {
	date := helper.DateFormatter(p.Date)
	device := &ShopspageDeviceBrand{
		Title:       p.Title,
		Date:        date,
		Description: p.Description,
		Image:       os.Getenv("ASSETS_URL") + p.Image,
	}

	if device.Title == "" || device.Date == "" {
		return nil, errors.New("shops page device primary fields can't be empty")
	}

	return device, nil
}
