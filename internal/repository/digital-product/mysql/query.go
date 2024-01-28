package repository

const (
	queryReduceStock = `
	UPDATE digital_products
	SET
		available_stock = available_stock - 1,
		updated_at = ?
	WHERE
		id = ?
	`
	queryGetDigitalProductsCredit = `
	SELECT
		dp.id AS ids,
		dp.ribbon_text1 AS ribbon_text1,
		dp.ribbon_text2 AS ribbon_text2,
		dp.ribbon_img1 AS ribbon_img1,
		dp.ribbon_img2 AS ribbon_img2,
		dp.business_pid,
		dp.total_stock,
		dp.available_stock,
		c.name_id,
		c.name_en,
		c.description_id,
		c.description_en,
		c.validity,
		c.price,
		c.original_price,
		dp.other_bonus,
		c.tc_id,
		c.tc_en,
		c.point,
		c.payment_dana,
		c.payment_shopeepay,
		c.payment_ovo,
		c.payment_gopay,
		c.payment_linkaja_wco,
		c.payment_linkaja_app
	FROM
		digital_products AS dp
	JOIN 
		digital_product_credits AS c
	ON
		c.digital_product_id = dp.id
	WHERE
	c.status = 1
	`

	queryFindDigitalProductsCredit = `
	SELECT
		dp.id AS ids,
		c.name_id,
		c.name_en,
		c.description_id,
		c.description_en,
		c.validity,
		c.price,
		c.original_price,
		dp.other_bonus,
		c.point,
		c.payment_dana,
		c.payment_shopeepay,
		c.payment_ovo,
		c.payment_gopay,
		c.payment_linkaja_wco,
		c.payment_linkaja_app
	FROM
		digital_products AS dp
	JOIN 
		digital_product_credits AS c
	ON
		c.digital_product_id = dp.id
	WHERE
	 	dp.id = ?  
	`
)
