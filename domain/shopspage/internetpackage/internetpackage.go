package domain

type ShopspageInternetPackageRequest struct {
	PackageType string `json:"package_type"`
}

type ShopspageInternetPackage struct {
	Postpaid []InternetPackage    `json:"postpaid,omitempty"`
	Prepaid  []InternetPackage    `json:"prepaid,omitempty"`
	CTA      InternetPackageCTA   `json:"cta"`
	Title    InternetPackageTitle `json:"title"`
}

type InternetPackage struct {
	PackageID     string `json:"id"`
	Title         string `json:"title"`
	TargetURL     string `json:"target_url"`
	Price         string `json:"price"`
	OriginalPrice string `json:"originalPrice"`
	Keyword       string `json:"string"`
	Category      string `json:"category"`
	Caption       string `json:"caption"`
	Image         string `json:"banner"`
}

type InternetPackageLabel struct {
	CTA   InternetPackageCTA
	TItle InternetPackageTitle
}

type InternetPackageCTA struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
	TargetURL string `json:"url"`
}

type InternetPackageTitle struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}

type ShopspageInternetPackageResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
