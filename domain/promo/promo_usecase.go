package domain

type PromoUseCase interface {
	GetPromoByTargetUrl(slug string) (*PromoResponse, error)
}
