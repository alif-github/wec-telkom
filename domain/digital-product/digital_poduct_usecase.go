package domain

type DigitalProductUseCase interface {
	GetCredit() (*CreditOffersV2, error)
	FindCredit(id string) (*ResponseFindCredit, error)
	ReduceStock(packageID string) error
}
