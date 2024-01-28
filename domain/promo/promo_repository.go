package domain

type Reader interface {
	GetPromoByTargetUrl(slug, date string, isLatest bool) (*Promo, error)
}

type PromoRepository interface {
	Reader
}
