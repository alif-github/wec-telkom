package repository

import (
	domain_listpoin "content-management/domain/listpoin"
	"content-management/helper"
	"content-management/helper/exceptions"
	"database/sql"
	"fmt"
	"strings"
)

type ListPoinRepository struct {
	db *sql.DB
}

func NewListPoinRepository(db *sql.DB) *ListPoinRepository {
	return &ListPoinRepository{db: db}
}

var (
	event = "INTERNAL|REPOSITORY|LISTPOIN|MYSQL|"
)

func (r *ListPoinRepository) GetListPoinByCategoryId(categoryId, subscriberType string) (*[]domain_listpoin.ListPoin, error) {
	var queryType string
	subscriberType = strings.ToLower(subscriberType)
	status := "1"
	fmt.Println("_____________________________________________")
	fmt.Println(subscriberType)
	if subscriberType == "prepaid" {
		queryType = "is_prepaid = ?"
	} else if subscriberType == "postpaid" {
		queryType = "is_postpaid = ?"
	}

	query := `SELECT title, keyword, poin, category, description, ribbon_text1, ribbon_text2, ribbon_image1, ribbon_image2, is_prepaid, is_postpaid FROM promo_poin WHERE parent_category = ? AND %s AND status = ?`
	query = fmt.Sprintf(query, queryType)
	rows, err := r.db.Query(query, categoryId, status, status)
	if err != nil {
		helper.StringLog("error", event+"GETLISTPOINBYCATEGORYID|QUERY Error: :"+err.Error())
		return nil, exceptions.ErrBadRequest
	}

	resListPoins := []domain_listpoin.ListPoin{}
	for rows.Next() {
		var (
			poin          domain_listpoin.ListPoin
			ribbon_text1  sql.NullString
			ribbon_text2  sql.NullString
			ribbon_image1 sql.NullString
			ribbon_image2 sql.NullString
		)
		err = rows.Scan(
			&poin.Title, &poin.Keyword, &poin.Poin, &poin.Category, &poin.Description, &ribbon_text1, &ribbon_text2, &ribbon_image1, &ribbon_image2, &poin.IsPrepaid, &poin.IsPostpaid,
		)

		if err != nil {
			helper.StringLog("error", event+"GETLISTPOINBYCATEGORYID|SCAN Error: :"+err.Error())
			return nil, exceptions.ErrBadRequest
		}

		if ribbon_text1.String != "" {
			poin.RibbonText = append(poin.RibbonText, ribbon_text1.String)
		}
		if ribbon_text2.String != "" {
			poin.RibbonText = append(poin.RibbonText, ribbon_text2.String)
		}
		if ribbon_image1.String != "" {
			poin.RibbonImage = append(poin.RibbonImage, ribbon_image1.String)
		}
		if ribbon_image2.String != "" {
			poin.RibbonImage = append(poin.RibbonImage, ribbon_image2.String)
		}

		if len(poin.RibbonText) == 0 {
			poin.RibbonText = make([]string, 0)
		}

		if len(poin.RibbonImage) == 0 {
			poin.RibbonImage = make([]string, 0)
		}

		listpoin, _ := domain_listpoin.GetListPoin(poin)
		resListPoins = append(resListPoins, *listpoin)
	}
	return &resListPoins, nil
}

func (r *ListPoinRepository) GetAllListPoin() (*[]domain_listpoin.ListPoin, error) {
	var poin domain_listpoin.ListPoin

	query := "SELECT title, keyword, poin, category, description, is_prepaid, is_postpaid FROM promo_poin"
	rows, err := r.db.Query(query)
	if err != nil {
		helper.StringLog("error", event+"GET_LIST_ALL_POIN|QUERY Error: :"+err.Error())
		return nil, err
	}

	resListPoins := []domain_listpoin.ListPoin{}
	for rows.Next() {
		err = rows.Scan(
			&poin.Title, &poin.Keyword, &poin.Poin, &poin.Category, &poin.Description, &poin.IsPrepaid, &poin.IsPostpaid,
		)
		if err != nil {
			helper.StringLog("error", event+"GET_LIST_ALL_POIN|SCAN Error: :"+err.Error())
			return nil, err
		}

		listpoin, _ := domain_listpoin.GetListPoin(poin)
		resListPoins = append(resListPoins, *listpoin)
	}

	return &resListPoins, nil
}

func (r *ListPoinRepository) CheckKeywordInMasterKeyword(keyword string) (bool, error) {
	var poin domain_listpoin.ListPoin
	var res bool

	query := "SELECT keyword FROM point_keyword WHERE keyword = ?"
	err := r.db.QueryRow(query, keyword).Scan(&poin.Keyword)
	if err != nil && err != sql.ErrNoRows {
		helper.StringLog("error", event+"GET_LIST_ALL_POIN|QUERY Error: :"+err.Error())
		return false, err
	}

	if poin.Keyword != "" {
		res = true
	} else {
		res = false
	}

	return res, nil
}

func (r *ListPoinRepository) DeleteKeyword(keyword string) error {
	query := "DELETE from promo_poin WHERE keyword = ?"
	_, err := r.db.Exec(query, keyword)
	if err != nil {
		helper.StringLog("error", event+"DELETE_KEYWORD|QUERY_ERROR: :"+err.Error())
		return err
	}

	return err
}
