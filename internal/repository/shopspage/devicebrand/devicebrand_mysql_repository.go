package repository

import (
	domain_shopsdevicebrand "content-management/domain/shopspage/devicebrand"
	"content-management/helper"
	"database/sql"
)

type ShopspageDeviceBrandRepository struct {
	db *sql.DB
}

func NewShopspageDeviceBrandRepository(db *sql.DB) *ShopspageDeviceBrandRepository {
	return &ShopspageDeviceBrandRepository{db: db}
}

func (r *ShopspageDeviceBrandRepository) GetShopspageDeviceBrands() (*[]domain_shopsdevicebrand.ShopspageDeviceBrand, error) {
	var devicebrand domain_shopsdevicebrand.ShopspageDeviceBrand
	status := "1"
	isDelete := "0"

	query := "SELECT title, COALESCE(description, ''), date, image FROM shops_device_brand WHERE status = ? AND is_delete = ? AND position IS NOT NULL ORDER BY -position DESC"
	rows, err := r.db.Query(query, status, isDelete)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/devicebrand] failed get list shops device brands :", err.Error())
		return nil, err
	}

	resShopsDeviceBrands := []domain_shopsdevicebrand.ShopspageDeviceBrand{}
	for rows.Next() {
		err = rows.Scan(
			&devicebrand.Title, &devicebrand.Description, &devicebrand.Date, &devicebrand.Image,
		)
		if err != nil {
			helper.StringLog("[internal/repository/shopspage/devicebrand/scan] error scan rows :", err.Error())
			return nil, err
		}

		shopsbanner, _ := domain_shopsdevicebrand.GetShopsBanner(devicebrand)
		resShopsDeviceBrands = append(resShopsDeviceBrands, *shopsbanner)
	}

	return &resShopsDeviceBrands, nil
}
