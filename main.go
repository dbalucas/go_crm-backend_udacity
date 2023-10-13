package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Customer struct for the database model
type Customer struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     int       `json:"phone,omitempty"`
	Contacted bool      `json:"contacted,omitempty"`
}

var mock_db []Customer
var CurrentID uuid.UUID

func init_db() {
	mock_db = append(mock_db, Customer{ID: getNextUniqueID(), Name: "Anton", Role: "account manager", Email: "sales.Anton@gmail.com", Phone: 4987654321, Contacted: true})
	mock_db = append(mock_db, Customer{ID: getNextUniqueID(), Name: "Bernd", Role: "marketing", Email: "marketing.bernd@gmail.com", Phone: 4987654321, Contacted: true})
	mock_db = append(mock_db, Customer{ID: getNextUniqueID(), Name: "Cäsar", Role: "product manager", Email: "product.caesar@web.com", Phone: 4987654321, Contacted: true})
	mock_db = append(mock_db, Customer{ID: getNextUniqueID(), Name: "Doris", Role: "admin", Email: "admin.doris@gmail.com", Phone: 4987654321, Contacted: true})
	mock_db = append(mock_db, Customer{ID: getNextUniqueID(), Name: "Emil", Role: "Evangelist", Email: "Evangelist.Emil@gmail.com", Phone: 4987654321, Contacted: true})
	fmt.Printf("Database initialized with %v entities. \n", len(mock_db))
}

func getNextUniqueID() uuid.UUID {
	CurrentID := uuid.New()
	return CurrentID
}

// TODO:
// brief overview of the API
func serverStatic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/html")
	http.ServeFile(w, r, "static/index.html")
}

// Recieve all customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mock_db)
}

// From POST request create new Customer if not exists
func addCustomer(w http.ResponseWriter, r *http.Request) {
	// new Entry var
	var newCustomer Customer
	// read the request body & parse JSON body
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // early  fail to avoid duplicate writing of status code
	}
	// insert new Entry into Customers map
	newCustomer.ID = getNextUniqueID()
	mock_db = append(mock_db, newCustomer)

	// output new entry
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCustomer)
}

// Retrieve single customer by ID in /customer/{id}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	CustomerID, _ := uuid.Parse(vars["id"])

	var foundCustomer *Customer
	for _, customer := range mock_db {
		if customer.ID == CustomerID {
			foundCustomer = &customer
			break
		}
	}

	if foundCustomer == nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	// if not found fail
	if false {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	const port = 3000
	const host = "localhost"

	var url = host + ":" + strconv.Itoa(port)

	init_db()

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
