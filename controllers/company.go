package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Solbaa/marvik/pkg/company"
	"github.com/Solbaa/marvik/viewmodels"
	"github.com/gorilla/mux"
)

type CompanyController struct {
	CompanyService company.Service
}

func NewCompanyService(c company.Service) *CompanyController {
	return &CompanyController{
		CompanyService: c,
	}
}

func (c *CompanyController) CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var company viewmodels.CompanyViewmodel
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		viewmodels.JSONError(w, viewmodels.ErrBadRequest, http.StatusBadRequest)
		return
	}

	modelCompany, err := c.CompanyService.CreateCompany(company)
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
		return
	}

	company = viewmodels.CompanyViewmodel{
		ID:   fmt.Sprint(modelCompany.ID),
		Name: company.Name,
	}
	json.NewEncoder(w).Encode(company)

}

func (pc *CompanyController) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var companyViewmodel []viewmodels.CompanyViewmodel
	companies, err := pc.CompanyService.GetAllCompanies()
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
		return
	}
	for _, company := range companies {
		comp := viewmodels.CompanyViewmodel{
			ID:   fmt.Sprint(company.ID),
			Name: company.Name,
		}
		companyViewmodel = append(companyViewmodel, comp)
	}

	json.NewEncoder(w).Encode(companyViewmodel)

}

func (pc *CompanyController) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	companyID := mux.Vars(r)["companyID"]
	company, err := pc.CompanyService.GetCompanyByID(companyID)
	if err != nil {
		viewmodels.JSONError(w, viewmodels.ErrCompanyNotFound, http.StatusNotFound)
		return
	}
	prodView := viewmodels.CompanyViewmodel{
		ID:   company.PublicID,
		Name: company.Name,
	}

	json.NewEncoder(w).Encode(prodView)

}
