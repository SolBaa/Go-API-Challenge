package viewmodels

import (
	"fmt"

	"github.com/Solbaa/marvik/models"
)

// /UserViewModel ths is the struct of the user we show to the client
type UserViewModel struct {
	ID       string             `json:"user_id,omitempty"`
	Name     string             `json:"name,omitempty"`
	LastName string             `json:"last_name,omitempty"`
	Email    string             `json:"email,omitempty"`
	Company  []CompanyViewmodel `json:"user_company"`
}

type CompanyViewmodel struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type CompanyRequest struct {
	Company models.Company `json:"user_companies"`
}

type CounterResponse struct {
	GetUsers         int `json:"get_users,omitempty"`
	GetUserByID      int `json:"get_user_by_id,omitempty"`
	AddCompanyToUser int `json:"add_company_to_user,omitempty"`
	DeleteUser       int `json:"delete_user,omitempty"`
	EndpointCounter  int `json:"endpoint_counter,omitempty"`
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
