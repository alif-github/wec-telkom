package repository

const (
	queryGetPromoByTargetUrl = `
		SELECT
			cp.id, cp.content_id, cp.category_id, cp.subcategory, cpe.content_promo_id, cp.type_promo, cp.menu_level,
			cp.startdate, cp.enddate, cp.page_title, cpe.page_title_en, cp.title, cpe.title_en,
			cp.slug, cp.button_text, cpe.button_text_en, cp.description, cpe.description_en,
			cp.hide_text, cpe.hide_text_en, cp.load_text, cpe.load_text_en, cp.offering_text,
			cpe.offering_text_en, cp.period_text, cpe.period_text_en, cp.terms_title,
			cpe.terms_title_en, cp.terms, cpe.terms_en, cp.status, cp.image, cp.created_at,
			COALESCE(cp.target_url, '') as target_url, COALESCE(cp.incomingdate, '') as incomingdate, COALESCE(cp.incoming_text, '') as incoming_text,
			COALESCE(cpe.incoming_text_en, '') as incoming_text_en, COALESCE(cp.tags_title, '') as tags_title,
			COALESCE(cpe.tags_title_en, '') as tags_title_en, COALESCE(cp.tags, '') as tags, COALESCE(cp.target_url, '') as target_url,
			COALESCE(ca.name, '') as category_name, COALESCE(sc.name, '') as subcategory_name
		FROM content_promo as cp
		JOIN content_promo_en as cpe
			ON cp.id = cpe.content_promo_id
		LEFT JOIN tag_category as ca
			ON cp.category_id = ca.id
		LEFT JOIN tag_subcategory as sc
			ON cp.subcategory = sc.id
		WHERE
			cp.target_url = %s
		AND
			cp.startdate <= %s
		ORDER BY
			cp.startdate DESC,
			cp.created_at ASC
		LIMIT
			1
	`
)
