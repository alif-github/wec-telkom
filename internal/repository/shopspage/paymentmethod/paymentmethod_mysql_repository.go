package repository

import (
	domain_shopspaymentmethod "content-management/domain/shopspage/paymentmethod"
	"content-management/helper"
	"database/sql"
)

type ShopspagePaymentMethodRepository struct {
	db *sql.DB
}

func NewShopspagePaymentMethodRepository(db *sql.DB) *ShopspagePaymentMethodRepository {
	return &ShopspagePaymentMethodRepository{db: db}
}

func (r *ShopspagePaymentMethodRepository) GetShopspagePaymentMethod() (*[]domain_shopspaymentmethod.ShopspagePaymentMethod, error) {
	var payment domain_shopspaymentmethod.ShopspagePaymentMethod
	status := "1"
	isDelete := "0"

	query := "SELECT title, COALESCE(description, ''), COALESCE(target_url, ''), date, image FROM shops_payment_method WHERE status = ? AND is_delete = ? AND position IS NOT NULL ORDER BY -position DESC"
	rows, err := r.db.Query(query, status, isDelete)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/paymentmethod] failed get list shops device brands :", err.Error())
		return nil, err
	}

	resShopsPayments := []domain_shopspaymentmethod.ShopspagePaymentMethod{}
	for rows.Next() {
		err = rows.Scan(
			&payment.Title, &payment.Description, &payment.TargetURL, &payment.Date, &payment.Image,
		)
		if err != nil {
			helper.StringLog("[internal/repository/shopspage/paymentmethod/scan] error scan rows :", err.Error())
			return nil, err
		}

		shopspayment, _ := domain_shopspaymentmethod.GetShopsPayments(payment)
		resShopsPayments = append(resShopsPayments, *shopspayment)
	}

	return &resShopsPayments, nil
}
