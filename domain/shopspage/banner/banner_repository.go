package domain

type Reader interface {
	GetShopspageBanners() (*[]ShopspageBanner, error)
	GetShopspageBannerLabel() (*ShopspageBannerLabel, error)
}

type ShopspageBannerRepository interface {
	Reader
}
