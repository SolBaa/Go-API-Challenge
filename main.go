package main

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/Solbaa/marvik/controllers"
	"github.com/Solbaa/marvik/models"

	"github.com/Solbaa/marvik/pkg/company"
	"github.com/Solbaa/marvik/pkg/db"
	"github.com/Solbaa/marvik/pkg/user"
	"github.com/gorilla/mux"
)

func main() {
	db := db.InitDb()
	err := db.AutoMigrate(&models.User{}, &models.Company{})
	if err != nil {
		fmt.Println(err)
	}
	companyService := company.NewService(db)
	userService := user.NewService(db, companyService)
	userController := controllers.NewUserService(userService)
	companyController := controllers.NewCompanyService(companyService)

	r := mux.NewRouter()
	{
		//Create a New User
		r.HandleFunc("/users", userController.CreateUser).Methods("POST")

		// Get a specific User
		r.HandleFunc("/users/{userID}", userController.GetOneUser).Methods("GET")

		//Get All the Users
		r.HandleFunc("/users", userController.GetAllUsers).Methods("GET")

		// Search user by query strings
		r.HandleFunc("/users-search", userController.SearchUsers).Methods("GET")
	}

	{
		//Create a company
		r.HandleFunc("/companies", companyController.CreateCompany).Methods("POST")
		//Get all available Companies
		r.HandleFunc("/companies", companyController.GetAllCompanies).Methods("GET")
		//Get company By ID
		r.HandleFunc("/companies/{companyID}", companyController.GetCompanyByID).Methods("GET")
		//Add a particular company to a particular User
		r.HandleFunc("/users/{userID}/{companyID}", userController.AddCompanyToUser).Methods("POST")
	}

	{
		//Delete a particular company in a particular User
		r.HandleFunc("/users/{userID}/companies/{companyID}", userController.DeleteCompanyFromUser).Methods("DELETE")

		// Delete all companies from a particular  User
		r.HandleFunc("/users/{userID}/companies", userController.DeleteAllCompaniesFromUser).Methods("DELETE")

		//Delete a particular User entirely
		r.HandleFunc("/users/{userID}", userController.DeleteUser).Methods("DELETE")
	}
	r.HandleFunc("/endpoint-count", userController.GetEndpointCount).Methods("GET")

	// r.PathPrefix("/swagger").Handler(http.StripPrefix("/swagger", http.FileServer(http.Dir("./swagger"))))

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
