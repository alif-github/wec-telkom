package domain

type ListPoinUseCase interface {
	GetListPoin(categoryName string) (*[]ListPoin, error)
}
