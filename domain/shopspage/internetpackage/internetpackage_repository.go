package domain

type Reader interface {
	GetShopspageInternetPackage(packageType string) (*[]InternetPackage, error)
	GetShopspageInternetPackageLabel(packageType string) (*InternetPackageLabel, error)
}

type ShopspageInternetPackageRepository interface {
	Reader
}
