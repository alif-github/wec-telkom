package repository

import (
	domain_shopsbanner "content-management/domain/shopspage/banner"
	"content-management/helper"
	"database/sql"
)

type ShopspageBannerRepository struct {
	db *sql.DB
}

func NewShopspageBannerRepository(db *sql.DB) *ShopspageBannerRepository {
	return &ShopspageBannerRepository{db: db}
}

func (r *ShopspageBannerRepository) GetShopspageBanners() (*[]domain_shopsbanner.ShopspageBanner, error) {
	var banner domain_shopsbanner.ShopspageBanner
	status := "1"
	isDelete := "0"

	query := "SELECT image, target_url FROM shops_banner WHERE status = ? AND is_delete = ? AND position IS NOT NULL ORDER BY -position DESC"
	rows, err := r.db.Query(query, status, isDelete)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/banner] failed get list shops page banner :", err.Error())
		return nil, err
	}

	resShopsBanners := []domain_shopsbanner.ShopspageBanner{}
	for rows.Next() {
		err = rows.Scan(
			&banner.Image, &banner.TargetURL,
		)
		if err != nil {
			helper.StringLog("[internal/repository/shopspage/banner/scan] error scan rows :", err.Error())
			return nil, err
		}

		shopsbanner, _ := domain_shopsbanner.GetShopsBanner(banner)
		resShopsBanners = append(resShopsBanners, *shopsbanner)
	}

	return &resShopsBanners, nil
}

func (r *ShopspageBannerRepository) GetShopspageBannerLabel() (*domain_shopsbanner.ShopspageBannerLabel, error) {
	var label domain_shopsbanner.ShopspageBannerLabel

	query := "SELECT text_button, text_button_en, COALESCE(target_url, '') FROM shops_banner_label LIMIT 1"
	err := r.db.QueryRow(query).Scan(&label.TextButton, &label.TextButtonEN, &label.TargetURL)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/bannerlabel] failed get banner label :", err.Error())
		return nil, err
	}

	return &label, nil
}
