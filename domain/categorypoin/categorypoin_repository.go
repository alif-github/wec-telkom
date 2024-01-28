package domain

type Reader interface {
	GetCategoryIdByName(categoryName string) (string, error)
}

type CategoryPoinRepository interface {
	Reader
}
