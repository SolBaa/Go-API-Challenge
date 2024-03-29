package company

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Solbaa/marvik/models"
	"github.com/Solbaa/marvik/viewmodels"
	"gorm.io/gorm"
)

type Service interface {
	CreateCompany(company viewmodels.CompanyViewmodel) (models.Company, error)
	GetAllCompanies() ([]models.Company, error)
	GetCompanyByID(CompanyID string) (models.Company, error)
}

type companyService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *companyService {
	return &companyService{
		db: db,
	}
}

func (pc *companyService) CreateCompany(company viewmodels.CompanyViewmodel) (models.Company, error) {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	companyID := y1.Intn(200)
	comp, err := pc.GetAllCompanies()

	for _, c := range comp {
		// Check if we already have a company with that name
		if c.Name == company.Name {
			fmt.Printf("Company already exists: %v", err)
			return models.Company{}, errors.New("Company already exists")
		}
	}
	companyModel := models.Company{
		Name:     company.Name,
		PublicID: strconv.Itoa(companyID),
	}

	err = pc.db.Omit("UserID").Create(&companyModel).Error
	if err != nil {
		return models.Company{}, err
	}
	return companyModel, nil
}

func (pc *companyService) GetAllCompanies() ([]models.Company, error) {
	companies := []models.Company{}
	err := pc.db.Omit("UserID").Find(&companies).Error
	if err != nil {
		return nil, err
	}

	return companies, nil

}

func (pc *companyService) GetCompanyByID(companyID string) (models.Company, error) {
	company := models.Company{}
	companyId, err := strconv.Atoi(companyID)
	if err != nil {
		return models.Company{}, err
	}
	err = pc.db.Omit("UserID").Where("id = ?", companyId).First(&company).Error
	if err != nil {
		return models.Company{}, err
	}

	return company, nil

}
