package main

import (
	"encoding/json"
	"errors"
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

// type CustomerRepository interface {
// 	// Save
// 	FindAll() ([]Customer, error)
// 	FindByID(uuid.UUID) (Customer, error)
// 	Save(Customer) (Customer, error)
// 	DeleteByID(uuid.UUID) error
// }

// var mockDb []Customer

type CustomerRepository struct {
	customers []Customer
}

func (c CustomerRepository) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func (c CustomerRepository) findIndexByID(id uuid.UUID) (int, error) {
	for index, customer := range c.customers {
		if customer.ID == id {
			return index, nil
		}
	}
	return 0, errors.New("not found")
}

func (c CustomerRepository) FindByID(id uuid.UUID) (Customer, error) {
	index, err := c.findIndexByID(id)
	if err != nil {
		return c.customers[index], nil
	}
	return Customer{}, errors.New("not found")
}

func (c CustomerRepository) Save(newCustomer Customer) (Customer, error) {
	c.customers = append(c.customers, newCustomer)
	return newCustomer, nil
}

func (c CustomerRepository) DeleteByID(id uuid.UUID) error {
	index, err := c.findIndexByID(id)
	if err != nil {
		c.customers = append(c.customers[:index], c.customers[index+1:]...)
		return nil
	}
	return errors.New("not found")

}

var CurrentID uuid.UUID

func init_db(customerRepo CustomerRepository) {
	customerRepo.Save(Customer{ID: getNextUniqueID(), Name: "Anton", Role: "account manager", Email: "sales.Anton@gmail.com", Phone: 4987654321, Contacted: true})
	customerRepo.Save(Customer{ID: getNextUniqueID(), Name: "Bernd", Role: "marketing", Email: "marketing.bernd@gmail.com", Phone: 4987654321, Contacted: true})
	customerRepo.Save(Customer{ID: getNextUniqueID(), Name: "Cäsar", Role: "product manager", Email: "product.caesar@web.com", Phone: 4987654321, Contacted: true})
	customerRepo.Save(Customer{ID: getNextUniqueID(), Name: "Doris", Role: "admin", Email: "admin.doris@gmail.com", Phone: 4987654321, Contacted: true})
	customerRepo.Save(Customer{ID: getNextUniqueID(), Name: "Emil", Role: "Evangelist", Email: "Evangelist.Emil@gmail.com", Phone: 4987654321, Contacted: true})
	fmt.Printf("Database initialized with %v entities. \n", len(customerRepo.customers))
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
	json.NewEncoder(w).Encode(mockDb)
	// add pagination
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

	// auslagern in eigenständige Funktion
	// insert new Entry into Customers map
	newCustomer.ID = getNextUniqueID()
	mockDb = append(mockDb, newCustomer)

	// output new entry
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCustomer)
}

// Retrieve single customer by ID in /customer/{id}
func getCustomer(w http.ResponseWriter, r *http.Request, c *CustomerRepository) {
	vars := mux.Vars(r)
	CustomerID, _ := uuid.Parse(vars["id"])

	var foundCustomer *Customer
	for _, customer := range mockDb {
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
	// Get ID from Var ULR
	vars := mux.Vars(r)
	customerID, _ := uuid.Parse(vars["id"])
	var foundCustomer *Customer
	var mockDbIndex int

	for sliceIndex, customer := range mockDb {
		if customer.ID == customerID {
			foundCustomer = &customer
			mockDbIndex = sliceIndex
			break
		} else {
			// do nothing
			foundCustomer = nil
		}
	}
	// if not found fail else remove
	if foundCustomer == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		mockDb = deleteAt(mockDb, mockDbIndex)
		fmt.Fprintf(w, "Custemer %v %v was removed", foundCustomer.ID, foundCustomer.Name)
	}
}

func deleteAt(slice []Customer, index int) []Customer {
	return append(slice[:index], slice[index+1:]...)
}

type MuxHandler = func(http.ResponseWriter, *http.Request)

type CustomHandler = func(http.ResponseWriter, *http.Request, *CustomerRepository)

func wrap(f CustomHandler, c *CustomerRepository) MuxHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, c)
	}
}

func main() {
	const port = 3000
	const host = "localhost"

	var url = host + ":" + strconv.Itoa(port)

	customerRepository := &CustomerRepository{}

	init_db(customerRepository)

	router := mux.NewRouter()

	router.HandleFunc("/", serverStatic)
	router.HandleFunc("/customers", wrap(getCustomer, customerRepository)).Methods("GET")
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
