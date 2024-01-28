package repository

import (
	domain_promo "content-management/domain/promo"
	"content-management/helper"
	"content-management/helper/exceptions"
	"database/sql"
	"strconv"
)

type PromoRepository struct {
	db *sql.DB
}

func NewPromoRepository(db *sql.DB) *PromoRepository {
	return &PromoRepository{db: db}
}

var (
	event = "INTERNAL|REPOSITORY|PROMO|PROMOREPOSITORY|"
)

func (r *PromoRepository) GetPromoByTargetUrl(targetUrl, date string, isLatest bool) (*domain_promo.Promo, error) {

	var promo domain_promo.Promo
	var status = 1
	var isDelete = 0
	var createdAt, target_url sql.NullString
	var contentPromoId, categoryId, subcategoryId, promoId sql.NullInt64
	queryGetPromo := `
        SELECT 
            cp.id, 
            cp.content_id, 
            cp.category_id, 
            cp.subcategory, 
            cpe.content_promo_id, 
            cp.type_promo, 
            cp.menu_level, 
            COALESCE(cp.startdate, '') as startdate, 
            COALESCE(cp.enddate, '') as enddate, 
            cp.page_title, 
            cpe.page_title_en, 
            cp.title, 
            cpe.title_en, 
            cp.slug, 
            cp.button_text, 
            cpe.button_text_en, 
            cp.description, 
            cpe.description_en, 
            cp.hide_text, 
            cpe.hide_text_en, 
            cp.load_text, 
            cpe.load_text_en, 
            cp.offering_text, 
            cpe.offering_text_en, 
            cp.period_text, 
            cpe.period_text_en, 
            cp.terms_title, 
            cpe.terms_title_en, 
            cp.terms, 
            cpe.terms_en, 
            cp.status, 
            cp.image, 
            cp.created_at, 
            COALESCE(cp.target_url, '') as target_url, 
            COALESCE(cp.incomingdate, '') as incomingdate, 
            COALESCE(cp.incoming_text, '') as incoming_text, 
            COALESCE(cpe.incoming_text_en, '') as incoming_text_en, 
            COALESCE(cp.tags_title, '') as tags_title, 
            COALESCE(cpe.tags_title_en, '') as tags_title_en, 
            COALESCE(cp.tags, '') as tags, 
            COALESCE(cp.target_url, '') as target_url, 
            COALESCE(ca.name, '') as category_name, 
            COALESCE(sc.name, '') as subcategory_name, 
            COALESCE(cp.using_time, '') as using_time, 
            COALESCE(cp.start_time, '') as start_time, 
            COALESCE(cp.end_time, '') as end_time 
        FROM 
            content_promo as cp 
        JOIN 
            content_promo_en as cpe 
        ON 
            cp.id = cpe.content_promo_id 
        LEFT JOIN 
            tag_category as ca 
        ON 
            cp.category_id = ca.id 
        LEFT JOIN 
            tag_subcategory as sc 
        ON 
            cp.subcategory = sc.id 
        WHERE 
            cp.target_url = ? 
        AND 
            cp.status = ? 
        AND 
            cp.is_delete = ? 
        ORDER BY 
            cp.updated_at DESC 
        LIMIT 1
    `

	rows, err := r.db.Query(queryGetPromo, targetUrl, status, isDelete)
	if err == sql.ErrNoRows {
		helper.StringLog("error", event+"GETPROMOBYTARGETURL|ERRNOROWS|TARGETURL:"+targetUrl+"|DATE:"+date+"|ISLATEST:"+strconv.FormatBool(isLatest)+" Error: "+err.Error())
		return nil, sql.ErrNoRows
	}

	if err != nil {
		helper.StringLog("error", event+"GETPROMOBYTARGETURL|QUERY TARGETURL:"+targetUrl+"|DATE:"+date+"|ISLATEST:"+strconv.FormatBool(isLatest)+" Error: "+err.Error())
		return nil, exceptions.ErrBadRequest
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&promoId, &promo.ContentId, &categoryId, &subcategoryId, &contentPromoId, &promo.TypePromo, &promo.ML,
			&promo.StartDate, &promo.EndDate, &promo.PageTitle, &promo.PageTitleEN, &promo.Title, &promo.TitleEN,
			&promo.Slug, &promo.ButtonText, &promo.ButtonTextEN, &promo.Description, &promo.DescriptionEN,
			&promo.HideText, &promo.HideTextEN, &promo.LoadText, &promo.LoadTextEN, &promo.OfferingText,
			&promo.OfferingTextEN, &promo.PeriodText, &promo.PeriodTextEN, &promo.TermsTitle,
			&promo.TermsTitleEN, &promo.Terms, &promo.TermsEN, &promo.Status, &promo.Image, &createdAt,
			&promo.TargetURL, &promo.IncomingDate, &promo.IncomingText, &promo.IncomingTextEN, &promo.TagsTitle,
			&promo.TagsTitleEN, &promo.Tags, &target_url,
			&promo.Category, &promo.Subcategory, &promo.UsingTime, &promo.StartTime, &promo.EndTime,
		); err != nil {
			helper.StringLog("error", event+"GETPROMOBYTARGETURL|SCAN TARGETURL:"+targetUrl+"|DATE:"+date+"|ISLATEST:"+strconv.FormatBool(isLatest)+" Error: "+err.Error())
			return nil, exceptions.ErrBadRequest
		}
	}

	return &promo, nil
}
