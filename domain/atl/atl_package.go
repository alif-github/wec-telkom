package domain

type ATLRequest struct {
	ProductName string   `json:"package_product_name,omitempty"`
	Validity    int      `json:"package_validity,omitempty"`
	StartPrice  int      `json:"package_start_price,omitempty"`
	EndPrice    int      `json:"package_end_price,omitempty"`
	RegionID    int      `json:"region_id,omitempty"`
	AtlUsed     string   `json:"atl_used,omitempty"`
	Category    []string `json:"package_category_id,omitempty"`
}

type ATLPackageResult struct {
	Result []ATLPackageResponse `json:"result"`
}
type ATLPackageResponse struct {
	ID                  string                `json:"id,omitempty"`
	BusinessPID         string                `json:"package_bid,omitempty"`
	CategoryID          int                   `json:"package_category_id,omitempty"`
	CategoryName        string                `json:"package_category,omitempty"`
	CategoryImage       string                `json:"package_category_image,omitempty"`
	CategoryColor       string                `json:"package_category_color,omiempty"`
	CategoryFontColor   string                `json:"package_category_font_color,omiempty"`
	ProductName         string                `json:"package_product_name,omitempty"`
	Disccount           int                   `json:"package_discount,omitempty"`
	Price               int                   `json:"package_price,omitempty"`
	BestPrice           int                   `json:"package_best_price,omitempty"`
	Validity            int                   `json:"package_validity,omitempty"`
	LangID              ProductInformation    `json:"ID,omitempty"`
	LangEN              ProductInformation    `json:"EN,omitempty"`
	PackageDistribution []PackageDistribution `json:"package_distribution,omitempty"`
}

type ProductInformation struct {
	Name    string `json:"package_name"`
	SubName string `json:"package_sub_name,omitempty"`
	Detail  string `json:"package_detail,omitempty"`
	TC      string `json:"package_tc,omittempty"`
}

type PackageDistribution struct {
	ID                        int    `json:"id"`
	PackageID                 string `json:"package_id"`
	PackageDistributionName   string `json:"package_distribution_name"`
	PackageDistributionAmount string `json:"package_distribution_amount"`
}

type TotalPackageDistribution struct {
	ID        int    `json:"id"`
	PackageID string `json:"package_id"`
}

type RegionResult struct {
	Result []RegionResponse `json:"result"`
}
type RegionResponse struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"region_name,omitempty"`
}
