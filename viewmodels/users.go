package viewmodels

import (
	"fmt"

	"github.com/Solbaa/marvik/models"
)

type UserViewModel struct {
	ID       string             `json:"user_id,omitempty"`
	Name     string             `json:"name,omitempty"`
	LastName string             `json:"last_name,omitempty"`
	Email    string             `json:"email"`
	Company  []CompanyViewmodel `json:"user_company"`
}

type CompanyViewmodel struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type CompanyRequest struct {
	Company models.Company `json:"user_companies"`
}

func ViewmodelToModel(companiesvm []CompanyViewmodel) []models.Company {
	companies := []models.Company{}
	for _, prod := range companiesvm {
		company := models.Company{
			PublicID: fmt.Sprint(prod.ID),
			Name:     prod.Name,
		}
		companies = append(companies, company)

	}
	return companies

}

func ModelToViewmodel(companies []models.Company) []CompanyViewmodel {
	var companiesModel []CompanyViewmodel
	for _, v := range companies {
		prodViewmodel := CompanyViewmodel{
			Name: v.Name,
		}
		companiesModel = append(companiesModel, prodViewmodel)
	}
	return companiesModel
}
