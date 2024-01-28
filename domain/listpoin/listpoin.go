package domain

import (
	"errors"
)

type ListPoin struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Poin        string   `json:"poin"`
	Keyword     string   `json:"keyword"`
	Category    string   `json:"category"`
	IsPrepaid   int      `json:"is_prepaid"`
	IsPostpaid  int      `json:"is_postpaid"`
	RibbonText  []string `json:"ribbon_text"`
	RibbonImage []string `json:"ribbon_image"`
}

type ListPoinResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"list_poin,omitempty"`
}

type ListPoinRequest struct {
	Category string `json:"category"`
	Msisdn   string `json:"msisdn"`
}

func NewListPoin(Title, Keyword, Poin, Category, Description string, IsPrepaid, IsPostpaid int) (*ListPoin, error) {

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

func GetListPoin(p ListPoin) (*ListPoin, error) {
	listpoin := &ListPoin{
		Title:       p.Title,
		Keyword:     p.Keyword,
		Poin:        p.Poin,
		Category:    p.Category,
		Description: p.Description,
		RibbonText:  p.RibbonText,
		RibbonImage: p.RibbonImage,
		IsPrepaid:   p.IsPrepaid,
		IsPostpaid:  p.IsPostpaid,
	}

	return listpoin, nil
}
