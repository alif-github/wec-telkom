package domain

type ShopspageBannerUsecase interface {
	GetShopspageBanners() (*ShopspageBannerLabelResponse, error)
}
