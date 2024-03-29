package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Solbaa/marvik/pkg/user"
	"github.com/Solbaa/marvik/viewmodels"

	"github.com/gorilla/mux"
)

type UserController struct {
	UserService user.Service
}

func NewUserService(u user.Service) *UserController {
	return &UserController{
		UserService: u,
	}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user viewmodels.UserViewModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		viewmodels.JSONError(w, viewmodels.ErrBadRequest, http.StatusBadRequest)
		return
	}

	modelsUser, err := c.UserService.CreateUser(user)
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
		return
	}
	company := viewmodels.ModelToViewmodel(modelsUser.Company)
	user = viewmodels.UserViewModel{
		ID:       fmt.Sprint(modelsUser.ID),
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Company:  company,
	}
	json.NewEncoder(w).Encode(user)

}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var usersViewmodel []viewmodels.UserViewModel
	users, err := c.UserService.GetAllUsers()
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
		return
	}
	for _, u := range users {
		comp := viewmodels.ModelToViewmodel(u.Company)
		user := viewmodels.UserViewModel{
			ID:       fmt.Sprint(u.ID),
			Name:     u.Name,
			LastName: u.LastName,
			Email:    u.Email,
			Company:  comp,
		}
		usersViewmodel = append(usersViewmodel, user)
	}

	json.NewEncoder(w).Encode(usersViewmodel)

}
func (c *UserController) SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	name := query.Get("name")
	lastName := query.Get("lastName")
	email := query.Get("email")
	company := query.Get("company")

	var usersViewmodel []viewmodels.UserViewModel
	users, err := c.UserService.GetAllWithFilter(name, lastName, email, company)
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
		return
	}
	for _, u := range users {
		comp := viewmodels.ModelToViewmodel(u.Company)
		user := viewmodels.UserViewModel{
			ID:       fmt.Sprint(u.ID),
			Name:     u.Name,
			LastName: u.LastName,
			Email:    u.Email,
			Company:  comp,
		}
		usersViewmodel = append(usersViewmodel, user)
	}

	json.NewEncoder(w).Encode(usersViewmodel)
}

func (c *UserController) GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["userID"]
	user, err := c.UserService.GetOneUser(userID)
	if err != nil {
		viewmodels.JSONError(w, viewmodels.ErrUserNotFound, http.StatusBadRequest)
		return
	}
	company := viewmodels.ModelToViewmodel(user.Company)
	uservm := viewmodels.UserViewModel{
		ID:       fmt.Sprint(user.ID),
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Company:  company,
	}
	json.NewEncoder(w).Encode(uservm)

}

func (c *UserController) AddCompanyToUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["userID"]
	companyID := mux.Vars(r)["companyID"]

	user, err := c.UserService.AddCompanyToUser(userID, companyID)
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
		return
	}
	company := viewmodels.ModelToViewmodel(user.Company)
	uservm := viewmodels.UserViewModel{
		ID:      user.PublicID,
		Company: company,
	}
	json.NewEncoder(w).Encode(uservm)

}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["userID"]

	user, err := c.UserService.DeleteUser(userID)
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
	}
	company := viewmodels.ModelToViewmodel(user.Company)
	uservm := viewmodels.UserViewModel{
		ID: fmt.Sprint(user.ID),

		Company: company,
	}
	err = json.NewEncoder(w).Encode(uservm)
	if err != nil {
		viewmodels.JSONError(w, viewmodels.ErrBadRequest, http.StatusBadRequest)
	}
}

func (c *UserController) DeleteAllCompaniesFromUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["userID"]
	user, err := c.UserService.DeleteAllCompaniesFromUser(userID)
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
	}
	company := viewmodels.ModelToViewmodel(user.Company)
	uservm := viewmodels.UserViewModel{
		ID:       user.PublicID,
		Company:  company,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
	err = json.NewEncoder(w).Encode(uservm)
	if err != nil {
		viewmodels.JSONError(w, viewmodels.ErrBadRequest, http.StatusBadRequest)
	}
}

func (c *UserController) DeleteCompanyFromUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["userID"]
	companyID := mux.Vars(r)["companyID"]
	comp, err := c.UserService.DeleteCompanyFromUser(userID, companyID)
	if err != nil {
		viewmodels.JSONError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(comp)

}

func (c *UserController) GetEndpointCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	counter, err := c.UserService.GetEndpointCount()
	if err != nil {
		viewmodels.JSONError(w, viewmodels.ErrInternal, http.StatusInternalServerError)
		return

	}
	resp := viewmodels.CounterResponse{
		GetUsers:         counter.GetUsers,
		GetUserByID:      counter.GetUserByID,
		AddCompanyToUser: counter.AddCompanyToUser,
		DeleteUser:       counter.DeleteUser,
		EndpointCounter:  counter.EndCounter,
	}
	json.NewEncoder(w).Encode(resp)
}
