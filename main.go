package main

import (
	"encoding/json"
	"fmt"
	"io"
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
var Customers = map[string]Customer{
	"1": {
		Name:  "Anton",
		Role:  "Sales",
		Email: "sales.Anton@gmail.com",
		Phone: "+49 876 54321",
	},
	"2": {
		Name:  "Bernd",
		Role:  "Sales",
		Email: "sales.bernd@gmail.com",
		Phone: "+49 876 54321",
	},
	"3": {
		Name:  "Cäsar",
		Role:  "Sales",
		Email: "sales.caesar@gmail.com",
		Phone: "+49 876 54321",
	},
	"4": {
		Name:  "Doris",
		Role:  "Sales",
		Email: "sales.doris@gmail.com",
		Phone: "+49 876 54321",
	},
	"5": {
		Name:  "Emil",
		Role:  "Sales",
		Email: "sales.Emil@gmail.com",
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

// From POST request create new Customer if not exists
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	// new Entry var
	var newCustomer Customer
	// read the request body
	requestBody, _ := io.ReadAll(r.Body)
	// parse JSON body
	json.Unmarshal(requestBody, &newCustomer)
	// insert new Entry into Customers map
	// for k, v := range newCustomer {
	// 	if _, ok := Customers[k]; !ok {
	// 		Customers[k] = v
	// 	}

	// }

	w.WriteHeader(http.StatusCreated)
}

// Retrieve single customer by ID in /customer/{id}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	id := mux.Vars(r)["id"]
	if _, ok := Customers[id]; ok {
		json.NewEncoder(w).Encode(Customers[id])
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

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
	err := http.ListenAndServe(url, router)
	if err != nil {
		panic(err)
	}
}

func main() {

	webServe()
}
