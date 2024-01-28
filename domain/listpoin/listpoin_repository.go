package domain

type Reader interface {
	GetListPoinByCategoryId(categoryId, subscriberType string) (*[]ListPoin, error)
	GetAllListPoin() (*[]ListPoin, error)
	CheckKeywordInMasterKeyword(keyword string) (bool, error)
	DeleteKeyword(keyword string) error
}

type ListPoinRepository interface {
	Reader
}
