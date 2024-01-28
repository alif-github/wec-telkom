package domain

import (
	"content-management/helper"
	"errors"
	"os"
)

type ShopspagePaymentMethod struct {
	Title       string `json:"title"`
	TargetURL   string `json:"target_url"`
	Date        string `json:"date"`
	Description string `json:"desc"`
	Image       string `json:"banner"`
}

type ShopspagePaymentMethodResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewShopspagePaymentMethod(title, date, description, image, target_url string) (*ShopspagePaymentMethod, error) {
	date = helper.DateFormatter(date)
	payment := &ShopspagePaymentMethod{
		Title:       title,
		Date:        date,
		Description: description,
		Image:       image,
		TargetURL:   target_url,
	}

	if payment.Title == "" && payment.Image == "" {
		return nil, errors.New("payment method primary fields can't be empty")
	}

	return payment, nil
}

func GetShopsPayments(p ShopspagePaymentMethod) (*ShopspagePaymentMethod, error) {
	date := helper.DateFormatter(p.Date)
	payment := &ShopspagePaymentMethod{
		Title:       p.Title,
		Date:        date,
		Description: p.Description,
		Image:       os.Getenv("ASSETS_URL") + p.Image,
		TargetURL:   p.TargetURL,
	}

	if payment.Title == "" || payment.Image == "" {
		return nil, errors.New("shops page payment primary fields can't be empty")
	}

	return payment, nil
}
