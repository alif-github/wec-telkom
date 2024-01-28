package repository

import (
	domain_shopsinternetpackage "content-management/domain/shopspage/internetpackage"
	"content-management/helper"
	"database/sql"
	"os"
	"strings"
)

type ShopspageInternetPackageRepository struct {
	db *sql.DB
}

func NewShopspageInternetPackageRepository(db *sql.DB) *ShopspageInternetPackageRepository {
	return &ShopspageInternetPackageRepository{db: db}
}

var (
	status   = "1"
	isDelete = "0"
)

func (r *ShopspageInternetPackageRepository) GetShopspageInternetPackage(packageType string) (*[]domain_shopsinternetpackage.InternetPackage, error) {
	var internet domain_shopsinternetpackage.InternetPackage
	packageType = strings.ToLower(packageType)

	query := "SELECT package_id, title, target_url, price, COALESCE(original_price, 0), keyword, category, COALESCE(caption, ''), image FROM shops_package WHERE package_type = ? AND status = ? AND is_delete = ? AND position IS NOT NULL ORDER BY -position DESC"
	rows, err := r.db.Query(query, packageType, status, isDelete)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/internetpackage] failed get list internet package :", err.Error())
		return nil, err
	}

	res := []domain_shopsinternetpackage.InternetPackage{}
	for rows.Next() {
		err = rows.Scan(
			&internet.PackageID, &internet.Title, &internet.TargetURL, &internet.Price,
			&internet.OriginalPrice, &internet.Keyword, &internet.Caption, &internet.Category, &internet.Image,
		)
		if err != nil {
			helper.StringLog("[internal/repository/shopspage/internetpackage/scan] error scan rows :", err.Error())
			return nil, err
		}

		internet.Image = os.Getenv("ASSETS_URL") + internet.Image
		res = append(res, internet)
	}

	return &res, nil
}

func (r *ShopspageInternetPackageRepository) GetShopspageInternetPackageLabel(packageType string) (*domain_shopsinternetpackage.InternetPackageLabel, error) {
	var cta domain_shopsinternetpackage.InternetPackageCTA
	var title domain_shopsinternetpackage.InternetPackageTitle
	packageType = strings.ToLower(packageType)

	query := "SELECT title, title_en, text_button, text_button_en, target_url FROM shops_package_label WHERE package_type = ? AND is_delete = ? AND status = ? LIMIT 1"
	err := r.db.QueryRow(query, packageType, isDelete, status).Scan(&title.Indonesia, &title.English,
		&cta.Indonesia, &cta.English, &cta.TargetURL)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/internetpackage/prepaidlabel] failed get bundling card label :", err.Error())
		return nil, err
	}

	res := &domain_shopsinternetpackage.InternetPackageLabel{
		CTA:   cta,
		TItle: title,
	}

	return res, nil
}
