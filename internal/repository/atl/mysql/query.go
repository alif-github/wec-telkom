package repository

const (
	queryGetPromoBySlug = `
		SELECT
			cp.id, cp.content_id, cp.category_id, cpe.content_promo_id, ca.name, cp.type_promo, cp.menu_level, 
			cp.startdate, cp.enddate, cp.page_title, cpe.page_title_en, cp.title, cpe.title_en, 
			cp.slug, cp.button_text, cpe.button_text_en, cp.description, cpe.description_en, 
			cp.hide_text, cpe.hide_text_en, cp.load_text, cpe.load_text_en, cp.offering_text, 
			cpe.offering_text_en, cp.period_text, cpe.period_text_en, cp.terms_title, 
			cpe.terms_title_en, cp.terms, cpe.terms_en, cp.status, cp.image, cp.created_at
		FROM content_promo as cp
		JOIN content_promo_en as cpe 
			ON cp.id = cpe.content_promo_id
		JOIN tag_category as ca
			ON cp.category_id = ca.id
		WHERE 
			cp.slug = %s
		AND
			cp.startdate >= %s
		ORDER BY
			cp.created_at
		LIMIT
			1
	`
	queryPromo = `
		SELECT
			cp.id, cp.content_id, cp.category_id, cpe.content_promo_id, ca.name, cp.type_promo, cp.menu_level, 
			cp.startdate, cp.enddate, cp.page_title, cpe.page_title_en, cp.title, cpe.title_en, 
			cp.slug, cp.button_text, cpe.button_text_en, cp.description, cpe.description_en, 
			cp.hide_text, cpe.hide_text_en, cp.load_text, cpe.load_text_en, cp.offering_text, 
			cpe.offering_text_en, cp.period_text, cpe.period_text_en, cp.terms_title, 
			cpe.terms_title_en, cp.terms, cpe.terms_en, cp.status, cp.image, cp.created_at
		FROM content_promo as cp
		JOIN content_promo_en as cpe 
			ON cp.id = cpe.content_promo_id
		JOIN tag_category as ca
			ON cp.category_id = ca.id
		WHERE 
			cp.slug = "%s"
		ORDER BY
			cp.created_at
		LIMIT
			1
	`
)
