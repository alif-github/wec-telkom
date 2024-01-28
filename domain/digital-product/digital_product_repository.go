package domain

type Reader interface {
	GetCredit() (*CreditOffersV2, error)
	FindCredit(id string) (*ResponseFindCredit, error)
	ReduceStock(packageID string) error
}

type DigitalProductRepository interface {
	Reader
}
