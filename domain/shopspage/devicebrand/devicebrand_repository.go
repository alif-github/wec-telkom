package domain

type Reader interface {
	GetShopspageDeviceBrands() (*[]ShopspageDeviceBrand, error)
}

type ShopspageDeviceBrandRepository interface {
	Reader
}
