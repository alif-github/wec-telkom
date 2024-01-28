package domain_test

import (
	domain_promo "content-management/domain/promo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPromo(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		mockPromo := domain_promo.Promo{
			ButtonText:     "",
			ButtonTextEN:   "",
			Category:       "Promo",
			Description:    "Promo Deskripsi",
			DescriptionEN:  "Promo Description",
			EndDate:        "29-07-2022",
			HideText:       "Sembunyikan",
			HideTextEN:     "Hide",
			Image:          "Gambar",
			LoadText:       "Selengkapnya",
			LoadTextEN:     "More",
			ML:             "ML-WEC-1233",
			OfferingText:   "Cek Penawaran Untuk Nomor Kamu",
			OfferingTextEN: "Check these offers for your number",
			PageTitle:      "Promo | TElkomsel",
			PeriodText:     "Periode Promo",
			PeriodTextEN:   "Promo Periode",
			Slug:           "nyobain-promo",
			StartDate:      "27-07-2022",
			Status:         "published",
			Title:          "Nyobain Promo",
			TermsTitle:     "Syarat dan Ketentuan",
			TermsTitleEN:   "Terms and Condition",
			Terms:          "Ini adalah syarat dan ketentuan",
			TermsEN:        "This is terms and condition",
			TypePromo:      "poin",
		}

		promo, err := domain_promo.NewPromo(
			mockPromo.ButtonText, mockPromo.ButtonTextEN, mockPromo.Category, mockPromo.Description, mockPromo.DescriptionEN,
			mockPromo.EndDate, mockPromo.HideText, mockPromo.HideTextEN, mockPromo.Image, mockPromo.LoadText, mockPromo.LoadTextEN,
			mockPromo.ML, mockPromo.OfferingText, mockPromo.OfferingTextEN, mockPromo.PageTitle, mockPromo.PeriodText,
			mockPromo.PeriodTextEN, mockPromo.Slug, mockPromo.StartDate, mockPromo.Status, mockPromo.Title,
			mockPromo.TermsTitle, mockPromo.TermsTitleEN, mockPromo.Terms, mockPromo.TermsEN, mockPromo.TypePromo,
		)

		assert.Nil(t, promo)
		assert.NotNil(t, err)
	})

	t.Run("empty", func(t *testing.T) {
		mockPromo := domain_promo.Promo{
			ButtonText:     "Cek Sekarang",
			ButtonTextEN:   "Check Now",
			Category:       "Promo",
			Description:    "Promo Deskripsi",
			DescriptionEN:  "Promo Description",
			EndDate:        "29-07-2022",
			HideText:       "Sembunyikan",
			HideTextEN:     "Hide",
			Image:          "Gambar",
			LoadText:       "Selengkapnya",
			LoadTextEN:     "More",
			ML:             "ML-WEC-1233",
			OfferingText:   "Cek Penawaran Untuk Nomor Kamu",
			OfferingTextEN: "Check these offers for your number",
			PageTitle:      "Promo | TElkomsel",
			PeriodText:     "Periode Promo",
			PeriodTextEN:   "Promo Periode",
			Slug:           "nyobain-promo",
			StartDate:      "27-07-2022",
			Status:         "published",
			Title:          "Nyobain Promo",
			TermsTitle:     "Syarat dan Ketentuan",
			TermsTitleEN:   "Terms and Condition",
			Terms:          "Ini adalah syarat dan ketentuan",
			TermsEN:        "This is terms and condition",
			TypePromo:      "poin",
		}

		promo, err := domain_promo.NewPromo(
			mockPromo.ButtonText, mockPromo.ButtonTextEN, mockPromo.Category, mockPromo.Description, mockPromo.DescriptionEN,
			mockPromo.EndDate, mockPromo.HideText, mockPromo.HideTextEN, mockPromo.Image, mockPromo.LoadText, mockPromo.LoadTextEN,
			mockPromo.ML, mockPromo.OfferingText, mockPromo.OfferingTextEN, mockPromo.PageTitle, mockPromo.PeriodText,
			mockPromo.PeriodTextEN, mockPromo.Slug, mockPromo.StartDate, mockPromo.Status, mockPromo.Title,
			mockPromo.TermsTitle, mockPromo.TermsTitleEN, mockPromo.Terms, mockPromo.TermsEN, mockPromo.TypePromo,
		)

		assert.Nil(t, err)
		assert.NotNil(t, promo)
	})
}
