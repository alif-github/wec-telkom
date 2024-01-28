package repository

import (
	domain "content-management/domain/atl"
	"content-management/helper"
	"database/sql"
	"fmt"
	"strings"
)

type ATLPackageRepository struct {
	db *sql.DB
}

func NewATLPackageRepository(db *sql.DB) *ATLPackageRepository {
	return &ATLPackageRepository{db: db}
}
func (r *ATLPackageRepository) GetRegion() (*domain.RegionResult, error) {
	var response domain.RegionResponse
	results := &domain.RegionResult{}
	query, err := r.db.Query("SELECT id,region_name from regions")
	if err != nil {

		helper.StringLog("error", err.Error())
		return results, err
	}

	defer query.Close()

	for query.Next() {
		if err := query.Scan(
			&response.ID, &response.Name,
		); err != nil {
			helper.StringLog("error", err.Error())
			return results, err
		}
		results.Result = append(results.Result, response)
	}

	return results, nil

}

func (r *ATLPackageRepository) GetPackage(request domain.ATLRequest) (*domain.ATLPackageResult, error) {
	var response domain.ATLPackageResponse
	var totalDistribution domain.TotalPackageDistribution

	results := &domain.ATLPackageResult{}
	var distribution domain.PackageDistribution
	var productInformationID domain.ProductInformation
	var productInformationEN domain.ProductInformation
	mapss := map[string]int{}
	category := strings.Join(request.Category, ",")
	// validity := strings.Join(request.Validity, ",")

	query := "SELECT p.id,p.package_bid,p.package_category_id,p.package_product_name,p.package_name_id,p.package_name_en,p.package_sub_name_id,p.package_sub_name_en,p.package_price,p.package_best_price, " +
		"p.package_validity,p.package_detail_id,p.package_detail_en,package_tc_id,package_tc_en, " +
		"pd.id,pd.package_id,pd.package_distribution_name ,pd.package_distribution_amount, " +
		"pc.package_categorie_name " +
		"FROM packages as p " +
		"JOIN package_categories as pc " +
		"ON pc.id = p.package_category_id " +
		"JOIN package_distributions as pd " +
		"ON p.id = pd.package_id " +
		"LEFT JOIN package_regions as pr " +
		"ON p.id = pr.package_id " +
		"WHERE " +
		"p.package_status = 1 " +
		"AND " +
		"p.package_product_name =  ? " +
		"AND " +
		"(pr.region_id = ? OR pr.region_id is null ) " +
		"AND " +
		"atl_used = ? "

	if len(category) != 0 {
		query += "AND " +
			"p.package_category_id IN ( " + category + ")"
	}
	if request.Validity == 1 {
		query += "AND " +
			"p.package_validity >= 1 " +
			"AND " +
			"p.package_validity < 7 "
	}
	if request.Validity == 7 {
		query += "AND " +
			"p.package_validity >= 7 " +
			"AND " +
			"p.package_validity < 30 "
	}
	if request.Validity == 30 {
		query += "AND " +
			"p.package_validity >= 30 "
	}
	if request.StartPrice != 0 && request.EndPrice != 0 {
		query += "AND " +
			"p.package_price BETWEEN " + fmt.Sprintf("%d", request.StartPrice) + " AND " + fmt.Sprintf("%d", request.EndPrice)
	}
	query += " order by package_category_id ASC "

	rows, err := r.db.Query(query, request.ProductName, request.RegionID, request.AtlUsed)

	if err != nil {

		helper.StringLog("error", err.Error())
		return results, err

	}
	defer rows.Close()

	queryDistribution, err := r.db.Query("SELECT COUNT(package_distributions.id) as id ,package_distributions.package_id FROM package_distributions Join packages ON packages.id = package_distributions.package_id GROUP BY packages.id")
	if err != nil {
		return results, err
	}
	defer queryDistribution.Close()

	for queryDistribution.Next() {
		if err := queryDistribution.Scan(
			&totalDistribution.ID, &totalDistribution.PackageID,
		); err != nil {
			helper.StringLog("error", err.Error())
			return results, err
		}
		mapss[totalDistribution.PackageID] = totalDistribution.ID
	}

	i := 0
	param := 0
	for rows.Next() {
		param++
		if err := rows.Scan(
			&response.ID, &response.BusinessPID, &response.CategoryID, &response.ProductName, &productInformationID.Name, &productInformationEN.Name, &productInformationID.SubName, &productInformationEN.SubName,
			&response.Price, &response.BestPrice, &response.Validity, &productInformationID.Detail, &productInformationEN.Detail, &productInformationID.TC, &productInformationEN.TC,
			&distribution.ID, &distribution.PackageID, &distribution.PackageDistributionName, &distribution.PackageDistributionAmount,
			&response.CategoryName,
		); err != nil {
			helper.StringLog("error", err.Error())
			return results, err
		}
		if response.CategoryID == 1 {
			response.CategoryImage = "https://www.telkomsel.com/sites/default/files/2022-09/Subtract2.png"
			response.CategoryColor = "#F2E2E2"
			response.CategoryFontColor = "linear-gradient(76.81deg, #B90024 15.71%, #FF0025 68.97%, #FD195E 94.61%)"
		}
		if response.CategoryID == 2 {
			response.CategoryImage = "https://www.telkomsel.com/sites/default/files/2022-09/Subtract.png"
			response.CategoryColor = "#EDF5FC"
			response.CategoryFontColor = "#0050AE"
		}
		if response.CategoryID == 3 {
			response.CategoryImage = "https://www.telkomsel.com/sites/default/files/2022-09/Sim-Card.png"
			response.CategoryColor = "#CDD6E7"
			response.CategoryFontColor = "#8E0D64"
		}
		if response.CategoryID == 4 {
			response.CategoryImage = "https://www.telkomsel.com/sites/default/files/2022-09/Roaming.png"
			response.CategoryColor = "#FFEED8"
			response.CategoryFontColor = "#FE6E00"

		}
		response.Disccount = 100 * (response.Price - response.BestPrice) / response.Price
		response.PackageDistribution = append(response.PackageDistribution, distribution)
		response.LangEN = productInformationEN
		response.LangID = productInformationID
		if mapss[distribution.PackageID] == param {
			results.Result = append(results.Result, response)
			response.PackageDistribution = nil

			param = 0
		}

		i++

	}

	return results, nil
}

func (r *ATLPackageRepository) FindPackage(id string) (*domain.ATLPackageResponse, error) {
	var response domain.ATLPackageResponse
	var totalDistribution domain.TotalPackageDistribution

	var distribution domain.PackageDistribution
	var productInformationID domain.ProductInformation
	var productInformationEN domain.ProductInformation
	mapss := map[string]int{}

	query := "SELECT p.id,p.package_bid,p.package_category_id,p.package_product_name,p.package_name_id,p.package_name_en,p.package_price,p.package_best_price, " +
		"p.package_validity,p.package_detail_id,p.package_detail_en,package_tc_id,package_tc_en, " +
		"pd.id,pd.package_id,pd.package_distribution_name ,pd.package_distribution_amount, " +
		"pc.package_categorie_name " +
		"FROM packages as p " +
		"JOIN package_categories as pc " +
		"ON pc.id = p.package_category_id " +
		"JOIN package_distributions as pd " +
		"ON p.id = pd.package_id " +
		"LEFT JOIN package_regions as pr " +
		"ON p.id = pr.package_id " +
		"WHERE " +
		"p.package_status = 1 " +
		"AND " +
		"p.id =  ? "

	query += " order by package_category_id DESC "

	rows, err := r.db.Query(query, id)

	if err != nil {

		helper.StringLog("error", err.Error())
		return &response, err

	}
	defer rows.Close()

	queryDistribution, err := r.db.Query("SELECT COUNT(package_distributions.id) as id ,package_distributions.package_id FROM package_distributions Join packages ON packages.id = package_distributions.package_id GROUP BY packages.id")
	if err != nil {
		return &response, err
	}
	defer queryDistribution.Close()

	for queryDistribution.Next() {
		if err := queryDistribution.Scan(
			&totalDistribution.ID, &totalDistribution.PackageID,
		); err != nil {
			helper.StringLog("error", err.Error())
			return &response, err
		}
		mapss[totalDistribution.PackageID] = totalDistribution.ID
	}

	param := 0
	for rows.Next() {
		param++
		if err := rows.Scan(
			&response.ID, &response.BusinessPID, &response.CategoryID, &response.ProductName, &productInformationID.Name, &productInformationEN.Name,
			&response.Price, &response.BestPrice, &response.Validity, &productInformationID.Detail, &productInformationEN.Detail, &productInformationID.TC, &productInformationEN.TC,
			&distribution.ID, &distribution.PackageID, &distribution.PackageDistributionName, &distribution.PackageDistributionAmount,
			&response.CategoryName,
		); err != nil {
			helper.StringLog("error", err.Error())
			return &response, err
		}

		response.Disccount = 100 * (response.Price - response.BestPrice) / response.Price
		response.PackageDistribution = append(response.PackageDistribution, distribution)
		response.LangEN = productInformationEN
		response.LangID = productInformationID
		if mapss[distribution.PackageID] == param {
			response.PackageDistribution = nil
			param = 0
		}
	}

	return &response, nil
}
