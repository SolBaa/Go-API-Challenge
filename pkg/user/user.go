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

var counter models.CounterEndpoints

type Service interface {
	CreateUser(User viewmodels.UserViewModel) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetAllWithFilter(name, lastName, email string) ([]models.User, error)
	GetOneUser(userID string) (models.User, error)
	AddCompanyToUser(userID, companyID string) (models.User, error)
	DeleteUser(userID string) (models.User, error)
	DeleteAllCompaniesFromUser(userID string) (models.User, error)
	DeleteCompanyFromUser(userID, companyID string) (models.User, error)
	GetEndpointCount() (*models.CounterEndpoints, error)
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
	counter.EndpointCounter("getUsers")

	User := []models.User{}

	err := c.db.Preload("Company").Find(&User).Error
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (c *userService) GetAllWithFilter(name, lastName, email string) ([]models.User, error) {
	user := []models.User{}
	if name != "" {
		err := c.db.Preload("Company").Where("name = ?", name).Find(&user).Error
		if err != nil {
			return nil, err
		}
	}
	if lastName != "" {
		err := c.db.Preload("Company").Where("last_name = ?", lastName).Find(&user).Error
		if err != nil {
			return nil, err
		}
	}
	if email != "" {
		err := c.db.Preload("Company").Where("email = ?", email).Find(&user).Error
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (c *userService) GetOneUser(userID string) (models.User, error) {
	counter.EndpointCounter("getUsersID")

	user := models.User{}
	userId, err := strconv.Atoi(userID)

	if err != nil {
		return models.User{}, err
	}

	err = c.db.Preload("Company").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (c *userService) AddCompanyToUser(userID, companyID string) (models.User, error) {
	counter.EndpointCounter("addCompany")
	userId, err := strconv.Atoi(userID)
	if err != nil {
		return models.User{}, err
	}

	comp, err := c.cs.GetCompanyByID(companyID)
	if err != nil {
		return models.User{}, err
	}
	newCompany := models.Company{
		UserID: uint(userId),
		Name:   comp.Name,
	}

	companies := []models.Company{}
	companies = append(companies, newCompany)
	userModel := models.User{
		Company: companies,
	}

	c.db.Model(&newCompany).Where("user_id=?", userModel.ID).Updates(&userModel).Save(&newCompany)

	return userModel, nil

}

func (c *userService) DeleteUser(userID string) (models.User, error) {
	counter.EndpointCounter("deleteUser")
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
	c.db.Model(&company).Where("user_id=?", userId).Delete(&company)

	return UserModel, nil

}

func (c *userService) DeleteCompanyFromUser(userID, companyID string) (models.User, error) {
	userId, err := strconv.Atoi(userID)
	if err != nil {
		return models.User{}, err
	}
	companyId, err := strconv.Atoi(companyID)
	if err != nil {
		return models.User{}, err
	}

	userModel := models.User{}
	err = c.db.Preload("Company").Where("id = ?", userId).Find(&userModel).Error
	if err != nil {
		fmt.Printf("User not found: %v", err)
		return models.User{}, err
	}
	company := models.Company{}
	err = c.db.Model(&company).Where("user_id=?", userId).Where("id = ?", companyId).Delete(&company).Error
	if err != nil {
		return models.User{}, err
	}

	return userModel, nil

}

func (c *userService) GetEndpointCount() (*models.CounterEndpoints, error) {
	endCounter := counter.EndpointCounter("endCounter")
	return endCounter, nil
}
