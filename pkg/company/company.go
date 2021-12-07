package company

import (
	"fmt"
	"strconv"

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
	// cs company.Service
}

func NewService(db *gorm.DB) *companyService {
	return &companyService{
		db: db,
	}
}

func (pc *companyService) CreateCompany(company viewmodels.CompanyViewmodel) (models.Company, error) {
	// x1 := rand.NewSource(time.Now().UnixNano())
	// y1 := rand.New(x1)
	// userID := y1.Intn(200)

	companyModel := models.Company{
		Name: company.Name,
		// PublicID: strconv.Itoa(userID),
	}
	err := pc.db.Create(&companyModel).Error
	if err != nil {
		return models.Company{}, err
	}
	return companyModel, nil
}

func (pc *companyService) GetAllCompanies() ([]models.Company, error) {
	companies := []models.Company{}
	err := pc.db.Find(&companies).Error
	if err != nil {
		return nil, err
	}

	return companies, nil

}

func (pc *companyService) GetCompanyByID(companyID string) (models.Company, error) {
	company := models.Company{}
	companyId, err := strconv.Atoi(companyID)
	fmt.Println("GetCompanyByID", companyId)
	if err != nil {
		return models.Company{}, err
	}
	err = pc.db.Where("id = ?", companyId).First(&company).Error
	if err != nil {
		return models.Company{}, err
	}

	return company, nil

}
