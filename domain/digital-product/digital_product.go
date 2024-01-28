package domain

type CreditOffersV2 struct {
	Result []ResultV2 `json:"result"`
}

type ResultV2 struct {
	ID            string          `json:"id"`
	BID           string          `json:"bid"`
	Name          string          `json:"name"`
	RibbonImage   []string        `json:"ribbon_image"`
	RibbonText    []string        `json:"ribbon_text"`
	ListPulsa     ListPulsaV2     `json:"list_pulsa"`
	PaymentMethod []PeymentMethod `json:"payment_method"`
}

type PeymentMethod struct {
	Amount     int    `json:"amount,omiempty"`
	MethodName string `json:"method_name,omiempty"`
}
type ListPulsaV2 struct {
	En Lang `json:"en"`
	ID Lang `json:"id"`
}

type Lang struct {
	AvailableStock     int           `json:"-"`
	Description        string        `json:"description"`
	Discount           interface{}   `json:"discount"`
	LabelValidity      string        `json:"label_validity"`
	Lang               string        `json:"lang"`
	Name               string        `json:"name"`
	OriginalPrice      int           `json:"original_price"`
	OtherBonus         string        `json:"other_bonus"`
	Price              int           `json:"price"`
	PurchasePercentage float32       `json:"purchase_percentage"`
	Subcategory        []Subcategory `json:"subcategory"`
	Tc                 string        `json:"tc"`
	TotalStock         int           `json:"-"`
	UseStock           bool          `json:"use_stock"`
}

type Subcategory struct {
	Advatage []Advatage `json:"advatage"`
}
type Advatage struct {
	Icon  string `json:"icon"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type PeymentMethodV2 struct {
	Dana       int `json:"dana"`
	ShopeePay  int `json:"shopeepay"`
	Gopay      int `json:"gopay"`
	Ovo        int `json:"ovo"`
	LinkAja    int `json:"linkaja_wco"`
	LinkAjaApp int `json:"linkaja_app"`
}

// type ResponseCreditOffersV2 struct {
// 	Result []ResponseFindCredit `json:"result"`
// }

type ResponseFindCredit struct {
	CreditId      string            `json:"CreditId"`
	ListPulsa2    []LangV2          `json:"ListCredit"`
	PaymentMethod []PeymentMethodV3 `json:"PaymentMethod"`
}

type AdvatageV2 struct {
	Icon  string `json:"Icon"`
	Label string `json:"Label"`
	Value string `json:"Value"`
}

type LangV2 struct {
	Advatage      []AdvatageV2 `json:"Advatage"`
	Description   string       `json:"Description"`
	Discount      interface{}  `json:"Discount"`
	LabelValidity string       `json:"LabelValidity"`
	Lang          string       `json:"Lang"`
	Name          string       `json:"Name"`
	OriginalPrice int          `json:"OriginalPrice"`
	OtherBonus    string       `json:"OtherBonus"`
	Price         int          `json:"Price"`
}
type PeymentMethodV3 struct {
	Amount     int    `json:"Amount,omiempty"`
	MethodName string `json:"MethodName,omiempty"`
}

type ListPulsaV3 struct {
	En LangV2 `json:"en"`
	ID LangV2 `json:"id"`
}
