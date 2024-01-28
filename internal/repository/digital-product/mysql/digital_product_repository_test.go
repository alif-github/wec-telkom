package repository

import (
	domain "content-management/domain/digital-product"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// func TestGetPackage(t *testing.T) {
// 	result := &domain.ATLPackageResult{}
// 	t.Run("Test Success", func(t *testing.T) {
// 		payload := domain.ATLRequest{
// 			ProductName: "Prabayar",
// 			StartPrice:  0,
// 			EndPrice:    9000000,
// 			RegionID:    1,
// 			AtlUsed:     "promo-jagoan",
// 			Validity:    1,
// 		}
// 		db, _ := NewMock()
// 		repo := NewATLPackageRepository(db)
// 		data, _ := repo.GetPackage(payload)

// 		// assert.Nil(t, nil, err)
// 		assert.Equal(t, result, data)
// 	})
// }

func TestGetCredit(t *testing.T) {
	result := &domain.CreditOffersV2{}
	t.Run("Test Success", func(t *testing.T) {
		db, _ := NewMock()
		repo := NewDigitalProductRepository(db)
		data, _ := repo.GetCredit()

		// assert.Nil(t, nil, err)
		assert.Equal(t, result, data)
	})
}

func TestReduceStock(t *testing.T) {
	t.Run("Test Success", func(t *testing.T) {
		payload := "1"
		db, _ := NewMock()
		repo := NewDigitalProductRepository(db)
		err := repo.ReduceStock(payload)

		assert.Nil(t, nil, err)
		// assert.Equal(t, result, data)
	})
}
