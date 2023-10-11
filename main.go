package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Customer struct for the database model
type Customer struct {
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted,omitempty"`
}

// Mock Database and init values
var Customers = map[uint32]Customer{
	1: {
		Name:  "Bernd",
		Role:  "Sales",
		Email: "sales.bernd@gmail.com",
		Phone: "+49 876 54321",
	},
	2: {
		Name:  "Bernd",
		Role:  "Sales",
		Email: "sales.bernd@gmail.com",
		Phone: "+49 876 54321",
	},
	3: {
		Name:  "Bernd",
		Role:  "Sales",
		Email: "sales.bernd@gmail.com",
		Phone: "+49 876 54321",
	},
	4: {
		Name:  "Bernd",
		Role:  "Sales",
		Email: "sales.bernd@gmail.com",
		Phone: "+49 876 54321",
	},
	5: {
		Name:  "Bernd",
		Role:  "Sales",
		Email: "sales.bernd@gmail.com",
		Phone: "+49 876 54321",
	},
}

// TODO:
// brief overview of the API
func serverStatic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/html")
	http.ServeFile(w, r, "static/index.html")
	w.WriteHeader(http.StatusOK)
}

// Recieve all customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(Customers)
	w.WriteHeader(http.StatusOK)
}
func addCustomer(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
}
func getCustomer(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
func updateCustomer(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// if not found fail
	if false {

		w.WriteHeader(http.StatusNotFound)
	} else {

		w.WriteHeader(http.StatusOK)
	}
}

func webServe() {
	const port = 3000
	const host = "localhost"

	var url = host + ":" + strconv.Itoa(port)

	router := mux.NewRouter()

	router.HandleFunc("/", serverStatic)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting the server on: ", url)
	http.ListenAndServe(url, router)
}

func main() {

	webServe()
}
