package repository

import (
	domainDP "content-management/domain/digital-product"
	"content-management/helper"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type DigitalProductRepository struct {
	db *sql.DB
}

func NewDigitalProductRepository(db *sql.DB) *DigitalProductRepository {
	return &DigitalProductRepository{db: db}
}

func (dpr *DigitalProductRepository) GetCredit() (*domainDP.CreditOffersV2, error) {
	var (
		creditOffers domainDP.ResultV2
		point        int
		ribbon_text1 string
		ribbon_text2 string
		ribbon_img1  string
		ribbon_img2  string
	)

	creditOffersV2 := domainDP.CreditOffersV2{}
	payment2 := domainDP.PeymentMethodV2{}
	price := 0
	query := fmt.Sprintf(queryGetDigitalProductsCredit)

	rows, err := dpr.db.Query(query)
	if err != nil {
		log.Println(err.Error())
		helper.StringLog("error", err.Error())
		return &creditOffersV2, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&creditOffers.ID, &ribbon_text1, &ribbon_text2, &ribbon_img1, &ribbon_img2, &creditOffers.BID, &creditOffers.ListPulsa.ID.TotalStock, &creditOffers.ListPulsa.ID.AvailableStock, &creditOffers.ListPulsa.ID.Name, &creditOffers.ListPulsa.En.Name,
			&creditOffers.ListPulsa.ID.Description, &creditOffers.ListPulsa.En.Description, &creditOffers.ListPulsa.ID.LabelValidity,
			&creditOffers.ListPulsa.ID.Price, &creditOffers.ListPulsa.ID.OriginalPrice, &creditOffers.ListPulsa.ID.OtherBonus,
			&creditOffers.ListPulsa.ID.Tc, &creditOffers.ListPulsa.En.Tc,
			&point,
			&payment2.Dana, &payment2.ShopeePay, &payment2.Ovo, &payment2.Gopay, &payment2.LinkAja, &payment2.LinkAjaApp,
		); err != nil {
			log.Println(err.Error())
			helper.StringLog("error", err.Error())
			return &creditOffersV2, err
		}

		percentage := (float32(creditOffers.ListPulsa.ID.TotalStock-creditOffers.ListPulsa.ID.AvailableStock) / float32(creditOffers.ListPulsa.ID.TotalStock)) * 100.00
		if percentage > 100 {
			percentage = 100
		}
		stock := os.Getenv("USE_STOCK")
		stockConv, err := strconv.ParseBool(stock)
		if err != nil {
			log.Println(err.Error())
		}

		if ribbon_img1 != "" {
			creditOffers.RibbonImage = append(creditOffers.RibbonImage, ribbon_img1)
		}
		if ribbon_img2 != "" {
			creditOffers.RibbonImage = append(creditOffers.RibbonImage, ribbon_img2)
		}

		if ribbon_text1 != "" {
			creditOffers.RibbonText = append(creditOffers.RibbonText, ribbon_text1)
		}
		if ribbon_text2 != "" {
			creditOffers.RibbonText = append(creditOffers.RibbonText, ribbon_text2)
		}

		creditOffers.ListPulsa.En.Lang = "EN"
		creditOffers.ListPulsa.ID.Lang = "ID"
		creditOffers.ListPulsa.En.UseStock = stockConv
		creditOffers.ListPulsa.ID.UseStock = stockConv
		creditOffers.ListPulsa.ID.PurchasePercentage = percentage
		creditOffers.ListPulsa.En.PurchasePercentage = percentage
		creditOffers.ListPulsa.En.TotalStock = creditOffers.ListPulsa.ID.TotalStock
		creditOffers.ListPulsa.En.AvailableStock = creditOffers.ListPulsa.ID.AvailableStock
		creditOffers.ListPulsa.En.OtherBonus = creditOffers.ListPulsa.ID.OtherBonus
		creditOffers.ListPulsa.En.Price = creditOffers.ListPulsa.ID.Price
		creditOffers.ListPulsa.En.OriginalPrice = creditOffers.ListPulsa.ID.OriginalPrice
		creditOffers.ListPulsa.En.LabelValidity = "+" + creditOffers.ListPulsa.ID.LabelValidity + " Days"
		creditOffers.ListPulsa.ID.LabelValidity = "+" + creditOffers.ListPulsa.ID.LabelValidity + " Hari"

		if creditOffers.ListPulsa.ID.Price < creditOffers.ListPulsa.ID.OriginalPrice {
			price = creditOffers.ListPulsa.ID.Price
		} else {
			price = creditOffers.ListPulsa.ID.OriginalPrice

		}
		if payment2.Dana == 1 {
			creditOffers.PaymentMethod = append(creditOffers.PaymentMethod, domainDP.PeymentMethod{Amount: price, MethodName: "DANA"})
		}
		if payment2.ShopeePay == 1 {
			creditOffers.PaymentMethod = append(creditOffers.PaymentMethod, domainDP.PeymentMethod{Amount: price, MethodName: "SHOPEEPAY"})
		}
		if payment2.Gopay == 1 {
			creditOffers.PaymentMethod = append(creditOffers.PaymentMethod, domainDP.PeymentMethod{Amount: price, MethodName: "GOPAY"})
		}
		if payment2.Ovo == 1 {
			creditOffers.PaymentMethod = append(creditOffers.PaymentMethod, domainDP.PeymentMethod{Amount: price, MethodName: "OVO"})

		}
		if payment2.LinkAja == 1 {

			creditOffers.PaymentMethod = append(creditOffers.PaymentMethod, domainDP.PeymentMethod{Amount: price, MethodName: "LINKAJA_WCO"})
		}
		if payment2.LinkAjaApp == 1 {

			creditOffers.PaymentMethod = append(creditOffers.PaymentMethod, domainDP.PeymentMethod{Amount: price, MethodName: "LINKAJA_APP"})
		}
		// Input Advantages into Sub Category
		var subCategory domainDP.Subcategory
		subCategory.Advatage = append(subCategory.Advatage, domainDP.Advatage{Icon: "time", Label: "Active Period", Value: creditOffers.ListPulsa.En.LabelValidity})
		subCategory.Advatage = append(subCategory.Advatage, domainDP.Advatage{Icon: "trophy", Label: "Poin", Value: fmt.Sprintf("+%d Point", point)})
		creditOffers.ListPulsa.En.Subcategory = append(creditOffers.ListPulsa.En.Subcategory, subCategory)
		subCategory.Advatage = nil
		subCategory.Advatage = append(subCategory.Advatage, domainDP.Advatage{Icon: "time", Label: "Masa Aktif", Value: creditOffers.ListPulsa.ID.LabelValidity})
		subCategory.Advatage = append(subCategory.Advatage, domainDP.Advatage{Icon: "trophy", Label: "Poin", Value: fmt.Sprintf("+%d Point", point)}) // salah di append, fix, tambahin TC juga di struct lalu query c.tc_id c.tc_en
		creditOffers.ListPulsa.ID.Subcategory = append(creditOffers.ListPulsa.ID.Subcategory, subCategory)
		subCategory.Advatage = nil

		creditOffersV2.Result = append(creditOffersV2.Result, creditOffers)
		creditOffers.PaymentMethod = nil
		creditOffers.ListPulsa.En.Subcategory = nil
		creditOffers.ListPulsa.ID.Subcategory = nil
		creditOffers.RibbonImage = nil
		creditOffers.RibbonText = nil

	}

	return &creditOffersV2, nil
}

func (dpr *DigitalProductRepository) FindCredit(id string) (*domainDP.ResponseFindCredit, error) {

	var responseFindCredit domainDP.ResponseFindCredit
	var listPulsaV3 domainDP.ListPulsaV3
	var point int

	payment2 := domainDP.PeymentMethodV2{}
	price := 0
	fmt.Println(id)
	query := fmt.Sprintf(queryFindDigitalProductsCredit)

	rows, err := dpr.db.Query(query, id)
	if err != nil {
		log.Println(err.Error())
		helper.StringLog("error", err.Error())
		return &responseFindCredit, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&responseFindCredit.CreditId, &listPulsaV3.ID.Name, &listPulsaV3.En.Name,
			&listPulsaV3.ID.Description, &listPulsaV3.En.Description, &listPulsaV3.ID.LabelValidity,
			&listPulsaV3.ID.Price, &listPulsaV3.ID.OriginalPrice, &listPulsaV3.ID.OtherBonus,
			&point,
			&payment2.Dana, &payment2.ShopeePay, &payment2.Ovo, &payment2.Gopay, &payment2.LinkAja, &payment2.LinkAjaApp,
		); err != nil {

			log.Println(err.Error())
			helper.StringLog("error", err.Error())
			return &responseFindCredit, err
		}

		listPulsaV3.En.Lang = "EN"
		listPulsaV3.ID.Lang = "ID"

		listPulsaV3.En.OtherBonus = listPulsaV3.ID.OtherBonus
		listPulsaV3.En.Price = listPulsaV3.ID.Price
		listPulsaV3.En.OriginalPrice = listPulsaV3.ID.OriginalPrice
		listPulsaV3.En.LabelValidity = "+" + listPulsaV3.ID.LabelValidity + " Days"
		listPulsaV3.ID.LabelValidity = "+" + listPulsaV3.ID.LabelValidity + " Hari"

		if listPulsaV3.ID.Price < listPulsaV3.ID.OriginalPrice {
			price = listPulsaV3.ID.Price
		} else {
			price = listPulsaV3.ID.OriginalPrice

		}

		if payment2.Dana == 1 {
			responseFindCredit.PaymentMethod = append(responseFindCredit.PaymentMethod, domainDP.PeymentMethodV3{Amount: price, MethodName: "DANA"})
		}
		if payment2.ShopeePay == 1 {
			responseFindCredit.PaymentMethod = append(responseFindCredit.PaymentMethod, domainDP.PeymentMethodV3{Amount: price, MethodName: "SHOPEEPAY"})
		}
		if payment2.Gopay == 1 {
			responseFindCredit.PaymentMethod = append(responseFindCredit.PaymentMethod, domainDP.PeymentMethodV3{Amount: price, MethodName: "GOPAY"})
		}
		if payment2.Ovo == 1 {
			responseFindCredit.PaymentMethod = append(responseFindCredit.PaymentMethod, domainDP.PeymentMethodV3{Amount: price, MethodName: "OVO"})

		}
		if payment2.LinkAja == 1 {

			responseFindCredit.PaymentMethod = append(responseFindCredit.PaymentMethod, domainDP.PeymentMethodV3{Amount: price, MethodName: "LINKAJA_WCO"})
		}
		if payment2.LinkAja == 1 {

			responseFindCredit.PaymentMethod = append(responseFindCredit.PaymentMethod, domainDP.PeymentMethodV3{Amount: price, MethodName: "LINKAJA_APP"})
		}

		listPulsaV3.ID.Advatage = append(listPulsaV3.ID.Advatage, domainDP.AdvatageV2{Icon: "time", Label: "Masa Aktif", Value: listPulsaV3.ID.LabelValidity})
		listPulsaV3.ID.Advatage = append(listPulsaV3.ID.Advatage, domainDP.AdvatageV2{Icon: "trophy", Label: "Poin", Value: fmt.Sprintf("+%d Point", point)})

		listPulsaV3.En.Advatage = append(listPulsaV3.En.Advatage, domainDP.AdvatageV2{Icon: "time", Label: "Active Period", Value: listPulsaV3.En.LabelValidity})
		listPulsaV3.En.Advatage = append(listPulsaV3.En.Advatage, domainDP.AdvatageV2{Icon: "trophy", Label: "Poin", Value: fmt.Sprintf("+%d Point", point)})

		responseFindCredit.ListPulsa2 = append(responseFindCredit.ListPulsa2, listPulsaV3.ID)
		responseFindCredit.ListPulsa2 = append(responseFindCredit.ListPulsa2, listPulsaV3.En)

	}

	return &responseFindCredit, nil
}

func (dpr *DigitalProductRepository) ReduceStock(packageID string) error {
	currentTime := time.Now().Local()
	_, err := dpr.db.Exec(queryReduceStock, currentTime, packageID)

	if err != nil {
		log.Println(err.Error())
		helper.StringLog("error", err.Error())
		return err
	}
	return err
}
