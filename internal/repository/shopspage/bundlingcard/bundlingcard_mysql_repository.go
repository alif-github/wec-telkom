package repository

import (
	domain_shopsbundlingcard "content-management/domain/shopspage/bundlingcard"
	"content-management/helper"
	"database/sql"
)

type ShopspageBundlingCardRepository struct {
	db *sql.DB
}

func NewShopspageBundlingCardRepository(db *sql.DB) *ShopspageBundlingCardRepository {
	return &ShopspageBundlingCardRepository{db: db}
}

func (r *ShopspageBundlingCardRepository) GetShopspageBundlingCard() (*[]domain_shopsbundlingcard.ShopspageBundlingCard, error) {
	var bundlingcard domain_shopsbundlingcard.ShopspageBundlingCard
	status := "1"
	isDelete := "0"

	query := "SELECT title, target_url, price, COALESCE(price_after_discount, ''), COALESCE(description, ''), date, image FROM shops_bundling_card WHERE status = ? AND is_delete = ? AND position IS NOT NULL ORDER BY -position DESC"
	rows, err := r.db.Query(query, status, isDelete)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/bundlingcard] failed get list shops bundling card :", err.Error())
		return nil, err
	}

	resShopsBundlingCard := []domain_shopsbundlingcard.ShopspageBundlingCard{}
	for rows.Next() {
		err = rows.Scan(
			&bundlingcard.Title, &bundlingcard.TargetURL, &bundlingcard.Price, &bundlingcard.PriceDiscount,
			&bundlingcard.Description, &bundlingcard.Date, &bundlingcard.Image,
		)
		if err != nil {
			helper.StringLog("[internal/repository/shopspage/bundlingcard/scan] error scan rows :", err.Error())
			return nil, err
		}

		bundling, _ := domain_shopsbundlingcard.GetShopsBundlingCard(bundlingcard)
		resShopsBundlingCard = append(resShopsBundlingCard, *bundling)
	}

	return &resShopsBundlingCard, nil
}

func (r *ShopspageBundlingCardRepository) GetShopspageBundlingCardLabel() (*domain_shopsbundlingcard.ShopspageBundlingCardCTA, error) {
	var label domain_shopsbundlingcard.ShopspageBundlingCardCTA
	var ctalabel domain_shopsbundlingcard.ShopspageBundlingCTALabel
	var ctatitle domain_shopsbundlingcard.ShopspageBundlingCTATitle

	query := "SELECT title, title_en, sub_title, text_button, text_button_en, description, COALESCE(target_url, '') FROM shops_bundling_card_label LIMIT 1"
	err := r.db.QueryRow(query).Scan(&ctatitle.Indonesia, &ctatitle.English, &label.Subtitle,
		&ctalabel.Indonesia, &ctalabel.English, &label.Description, &label.TargetURL)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/bundlingcard/label] failed get bundling card label :", err.Error())
		return nil, err
	}

	res, _ := domain_shopsbundlingcard.NewShopspageBundlingCardCTA(label.Description, label.Subtitle, label.TargetURL, ctalabel, ctatitle)

	return res, nil
}
