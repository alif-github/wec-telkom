package repository

import (
	"content-management/helper"
	"database/sql"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

var (
	event = "INTERNAL|REPOSITORY|CATEGORYPOIN|MYSQL|"
)

func (r *CategoryRepository) GetCategoryIdByName(categoryName string) (string, error) {
	var categoryId sql.NullString

	query := `
		SELECT  
			id 
		FROM 
			tag_category 
		WHERE 
			name = ?
	`

	err := r.db.QueryRow(query, categoryName).Scan(&categoryId)
	if err != nil {
		helper.StringLog("error", event+"GETCATEGORYIDBYNAME Error "+err.Error())
		return "", err
	}

	return categoryId.String, nil
}
