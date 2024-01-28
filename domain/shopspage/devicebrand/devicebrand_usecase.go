package domain

type ShopspageDeviceBrandUsecase interface {
	GetShopspageDeviceBrands() (*[]ShopspageDeviceBrand, error)
}
