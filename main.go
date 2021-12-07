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
	//Create a New User
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	// Get a specific User
	r.HandleFunc("/users/{userID}", userController.GetOneUser).Methods("GET")
	//Get All the Users
	r.HandleFunc("/users", userController.GetAllUsers).Methods("GET")

	r.HandleFunc("/companies", companyController.CreateCompany).Methods("POST")
	//Get all available Companies
	r.HandleFunc("/companies", companyController.GetAllCompanies).Methods("GET")
	// //Get Product By ID
	// r.HandleFunc("/Companies/{productID}", companyController.GetCompanyByID).Methods("GET")
	// //Add a particular Product to a particular Shopping User
	// r.HandleFunc("/User/{UserID}", userController.AddCompanyToUser).Methods("POST")
	// // Modify the amount of a particular product in a particular Shopping User
	// // r.HandleFunc("/User/{UserID}/Companies/{productID}", userController.ModifyProductAmount).Methods("PUT")
	// //Delete a particular product in a particular Shopping User
	// r.HandleFunc("/User/{UserID}/Companies/{productID}", userController.DeleteCompanyFromUser).Methods("DELETE")
	// // Delete all Companies from a particular Shopping User
	// r.HandleFunc("/User/{UserID}/Companies", userController.DeleteAllCompaniesFromUser).Methods("DELETE")
	// //Delete a particular Shopping User entirely
	// r.HandleFunc("/User/{UserID}", userController.DeleteUser).Methods("DELETE")
	r.PathPrefix("/swagger").Handler(http.StripPrefix("/swagger", http.FileServer(http.Dir("./swagger"))))

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
