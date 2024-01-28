package domain

type Reader interface {
	GetShopspagePaymentMethod() (*[]ShopspagePaymentMethod, error)
}

type ShopspagePaymentMethodRepository interface {
	Reader
}
