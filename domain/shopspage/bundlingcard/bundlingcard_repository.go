package domain

type Reader interface {
	GetShopspageBundlingCard() (*[]ShopspageBundlingCard, error)
	GetShopspageBundlingCardLabel() (*ShopspageBundlingCardCTA, error)
}

type ShopspageBundlingCardRepository interface {
	Reader
}
