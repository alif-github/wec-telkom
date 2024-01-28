package domain

type Reader interface {
	GetShopspageExpressInput() (*[]ShopspageExpressInput, error)
}

type ShopspageExpressInputRepository interface {
	Reader
}
