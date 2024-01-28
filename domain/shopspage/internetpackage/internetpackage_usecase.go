package domain

type ShopspageInternetPackageUsecase interface {
	GetShopspageInternetPackage(packageType string) (*ShopspageInternetPackage, error)
}
