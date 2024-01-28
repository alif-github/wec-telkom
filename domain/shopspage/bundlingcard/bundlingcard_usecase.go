package domain

type ShopspageBundlingCardUsecase interface {
	GetShopspageBundlingCard() (*ShopspageBundlingCTAResponse, error)
}
