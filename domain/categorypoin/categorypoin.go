package domain

import (
	"errors"
)

type ListPoin struct {
	Title       string
	Keyword     string
	Poin        string
	Category    string
	Description string
	IsPrepaid   bool
	IsPostpaid  bool
}

type ListPoinResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ListPoinRequest struct {
	Category string `json:"category"`
}

func NewListPoin(Title, Keyword, Poin, Category, Description string, IsPrepaid, IsPostpaid bool) (*ListPoin, error) {

	listPoin := &ListPoin{
		Title:       Title,
		Keyword:     Keyword,
		Poin:        Poin,
		Category:    Category,
		Description: Description,
		IsPrepaid:   IsPrepaid,
		IsPostpaid:  IsPostpaid,
	}

	if listPoin.Title == "" || listPoin.Keyword == "" || listPoin.Poin == "" || listPoin.Category == "" || listPoin.Description == "" {
		return nil, errors.New("primary fields can't be empty")
	}

	return listPoin, nil
}
