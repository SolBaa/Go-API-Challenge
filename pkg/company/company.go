package company

import (
	"github.com/Solbaa/marvik/models"
	"github.com/Solbaa/marvik/viewmodels"
	"gorm.io/gorm"
)

type Service interface {
	CreateCompany(company viewmodels.CompanyViewmodel) (models.Company, error)
	GetAllCompanies() ([]models.Company, error)
	GetCompanyByID(id string) (models.Company, error)
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
	return []models.Company{}, nil

}

func (pc *companyService) GetCompanyByID(id string) (models.Company, error) {
	return models.Company{}, nil

}
