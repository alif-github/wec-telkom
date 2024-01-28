package repository

import (
	domain_shopsexpressinput "content-management/domain/shopspage/expressinput"
	"content-management/helper"
	"database/sql"
)

type ShopspageExpressInputRepository struct {
	db *sql.DB
}

func NewShopspageExpressInputRepository(db *sql.DB) *ShopspageExpressInputRepository {
	return &ShopspageExpressInputRepository{db: db}
}

func (r *ShopspageExpressInputRepository) GetShopspageExpressInput() (*[]domain_shopsexpressinput.ShopspageExpressInput, error) {
	var input domain_shopsexpressinput.ShopspageExpressInput
	var title domain_shopsexpressinput.ShopspageExpressInputTitle

	status := "1"
	isDelete := "0"

	query := "SELECT title, title_en, COALESCE(description, ''), date, target_url, image FROM shops_express_input WHERE status = ? AND is_delete = ? AND position IS NOT NULL ORDER BY -position DESC"
	rows, err := r.db.Query(query, status, isDelete)
	if err != nil {
		helper.StringLog("[internal/repository/shopspage/expressinput] failed get list shops device brands :", err.Error())
		return nil, err
	}

	resExpressInputs := []domain_shopsexpressinput.ShopspageExpressInput{}
	for rows.Next() {
		err = rows.Scan(
			&title.Indonesia, &title.English, &input.Description, &input.Date, &input.TargetURL, &input.Image,
		)
		if err != nil {
			helper.StringLog("[internal/repository/shopspage/expressinput/scan] error scan rows :", err.Error())
			return nil, err
		}

		expressInput, _ := domain_shopsexpressinput.GetShopsExpressInput(input)
		inputTitle, _ := domain_shopsexpressinput.GetShopsExpressInputTitle(title)
		input, _ := domain_shopsexpressinput.NewShopspageExpressInput(expressInput.Image, expressInput.Description, expressInput.Date, expressInput.TargetURL, *inputTitle)
		resExpressInputs = append(resExpressInputs, *input)
	}

	return &resExpressInputs, nil
}
