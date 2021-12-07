package user

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Solbaa/marvik/models"
	"github.com/Solbaa/marvik/pkg/company"
	"github.com/Solbaa/marvik/viewmodels"
	"gorm.io/gorm"
)

type Service interface {
	CreateUser(User viewmodels.UserViewModel) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetAllWithFilter(name, lastName, email string) ([]models.User, error)
	GetOneUser(userID string) (models.User, error)
	AddCompanyToUser(userID, companyID string, user viewmodels.CompanyRequest) (models.User, error)
	// ModifyProductAmount(userID, productID string, product viewmodels.ProductRequest) error
	DeleteUser(userID string) (models.User, error)
	DeleteAllCompaniesFromUser(userID string) (models.User, error)
	DeleteCompanyFromUser(userID, companyID string) (models.User, error)
}

type userService struct {
	db *gorm.DB
	cs company.Service
}

func NewService(db *gorm.DB, cs company.Service) *userService {
	return &userService{
		db: db,
		cs: cs,
	}
}

func (c *userService) CreateUser(User viewmodels.UserViewModel) (models.User, error) {
	company := viewmodels.ViewmodelToModel(User.Company)
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	userID := y1.Intn(200)

	users, err := c.GetAllUsers()
	userModel := models.User{}
	for _, u := range users {
		if u.Email == User.Email {
			fmt.Printf("User already exists: %v", err)
			return models.User{}, errors.New("User already exists")
		}

		userModel = models.User{
			PublicID: strconv.Itoa(userID),
			Company:  company,
			Name:     User.Name,
			LastName: User.LastName,
			Email:    User.Email,
		}
	}

	err = c.db.Create(&userModel).Error
	if err != nil {
		return models.User{}, err
	}
	return userModel, nil

}

func (c *userService) GetAllUsers() ([]models.User, error) {
	User := []models.User{}

	err := c.db.Preload("Company").Find(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}
func (c *userService) GetAllWithFilter(name, lastName, email string) ([]models.User, error) {
	User := []models.User{}
	if name != "" {
		err := c.db.Preload("Company").Where("name = ?", name).Find(&User).Error
		if err != nil {
			return nil, err
		}
	}
	if lastName != "" {
		err := c.db.Preload("Company").Where("last_name = ?", lastName).Find(&User).Error
		if err != nil {
			return nil, err
		}
	}
	if email != "" {
		err := c.db.Preload("Company").Where("email = ?", email).Find(&User).Error
		if err != nil {
			return nil, err
		}
	}

	return User, nil
}
func (c *userService) GetOneUser(userID string) (models.User, error) {
	user := models.User{}
	userId, err := strconv.Atoi(userID)
	fmt.Printf("USERRRERERERER %v\n", userId)
	if err != nil {
		return models.User{}, err
	}

	err = c.db.Preload("Company").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (c *userService) AddCompanyToUser(userID, companyID string, company viewmodels.CompanyRequest) (models.User, error) {
	userId, err := strconv.Atoi(userID)
	if err != nil {
		return models.User{}, err
	}
	UserModel, err := c.GetOneUser(userID)
	if err != nil {
		return models.User{}, err
	}
	fmt.Printf("USSSEERRRR %+v", UserModel)

	prod, err := c.cs.GetCompanyByID(companyID)
	if err != nil {
		return models.User{}, err
	}
	newCompany := models.Company{
		Code:     userID,
		PublicID: prod.PublicID,
		Name:     prod.Name,
	}

	companies := []models.Company{}

	for _, item := range UserModel.Company {
		if item.PublicID == prod.PublicID {
			fmt.Println("Company already in User")
			return models.User{}, err
		}

	}
	companies = append(companies, newCompany)
	UserModel.Company = companies

	c.db.Model(&newCompany).Where("id = ?", userId).Updates(&UserModel).Save(&newCompany)
	fmt.Printf("USSSEERRRR %+v", UserModel)
	return UserModel, nil

}

func (c *userService) DeleteUser(userID string) (models.User, error) {
	userId, err := strconv.Atoi(userID)
	if err != nil {
		return models.User{}, err
	}
	UserModel, err := c.GetOneUser(userID)
	if err != nil {
		fmt.Printf("User not found: %v", err)
		return models.User{}, err
	}

	c.db.Where("id = ?", userId).Delete(&UserModel)
	return UserModel, nil

}

func (c *userService) DeleteAllCompaniesFromUser(userID string) (models.User, error) {
	userId, err := strconv.Atoi(userID)
	if err != nil {
		return models.User{}, err
	}
	UserModel, err := c.GetOneUser(userID)
	if err != nil {
		fmt.Printf("User not found: %v", err)
		return models.User{}, err
	}
	company := models.Company{}
	c.db.Model(&company).Where("User_id=?", userId).Delete(&company)

	return UserModel, nil

}

func (c *userService) DeleteCompanyFromUser(userID, companyID string) (models.User, error) {
	userId, err := strconv.Atoi(userID)
	if err != nil {
		return models.User{}, err
	}
	UserModel := models.User{}
	err = c.db.Preload("company").Where("id = ?", userId).Find(&UserModel).Error
	if err != nil {
		fmt.Printf("User not found: %v", err)
		return models.User{}, err
	}
	company := models.Company{}
	err = c.db.Model(&company).Where("User_id=?", userId).Where("public_id = ?", companyID).Delete(&company).Error
	if err != nil {
		return models.User{}, err
	}

	return UserModel, nil

}
