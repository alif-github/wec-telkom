package domain

type ListPoinUseCase interface {
	GetListPoin(categoryName, msisdn string) (*[]ListPoin, error)
	DeleteKeywordAutomate() error
}
