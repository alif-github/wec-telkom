package domain

type ShopspagePaymentMethodUsecase interface {
	GetShopspagePaymentMethod() (*[]ShopspagePaymentMethod, error)
}
