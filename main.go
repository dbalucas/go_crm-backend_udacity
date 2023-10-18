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

// Repository alice of all Customers
type CustomerRepository struct {
	customers []Customer
}

// type CustomerRepository interface {
// 	// Save
// 	FindAll() ([]Customer, error)
// 	FindByID(uuid.UUID) (Customer, error)
// 	Save(Customer) (Customer, error)
// 	DeleteByID(uuid.UUID) error
// }

// retrieve all customers from database
func (c *CustomerRepository) FindAll() ([]Customer, error) {
	return c.customers, nil
}

// retrieve the database index by uuid
func (c *CustomerRepository) findIndexByID(id uuid.UUID) (int, error) {
	for index, customer := range c.customers {
		if customer.ID == id {
			return index, nil
		}
	}
	return 0, errors.New("not found")
}

// find a user by uuid in database if exists
func (c *CustomerRepository) FindByID(id uuid.UUID) (Customer, error) {
	index, err := c.findIndexByID(id)
	if err != nil {
		return c.customers[index], nil
	}
	return Customer{}, errors.New("not found")
}

// Save newCustomer to Database (UUID generated automatically)
func (c *CustomerRepository) Save(newCustomer Customer) (Customer, error) {
	// create new ID if not exists
	if newCustomer.ID == uuid.Nil {
		newCustomer.ID = getNextUniqueID()
	}
	c.customers = append(c.customers, newCustomer)
	return newCustomer, nil
}

// delete a customer by uuid from database
func (c *CustomerRepository) DeleteByID(id uuid.UUID) error {
	index, err := c.findIndexByID(id)
	if err != nil {
		return errors.New("can not delete by ID: " + id.String() + " not found")
	}
	c.customers = deleteAtIndex(c.customers, index)
	return nil

}

// from inserted slice, remove item by slice index
func deleteAtIndex(slice []Customer, index int) []Customer {
	return append(slice[:index], slice[index+1:]...)
}

func init_db(customerRepository *CustomerRepository) {
	customerRepository.Save(Customer{Name: "Anton", Role: "account manager", Email: "sales.Anton@gmail.com", Phone: 4987654321, Contacted: true})
	customerRepository.Save(Customer{Name: "Bernd", Role: "marketing", Email: "marketing.bernd@gmail.com", Phone: 4987654321, Contacted: true})
	customerRepository.Save(Customer{Name: "Cäsar", Role: "product manager", Email: "product.caesar@web.com", Phone: 4987654321, Contacted: true})
	customerRepository.Save(Customer{Name: "Doris", Role: "admin", Email: "admin.doris@gmail.com", Phone: 4987654321, Contacted: true})
	customerRepository.Save(Customer{Name: "Emil", Role: "Evangelist", Email: "Evangelist.Emil@gmail.com", Phone: 4987654321, Contacted: true})
	fmt.Printf("Database initialized with %v entities. \n", len(customerRepository.customers))
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
func getCustomers(w http.ResponseWriter, r *http.Request, c *CustomerRepository) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c.customers)
	// add pagination
}

// From POST request create new Customer if not exists
func addCustomer(w http.ResponseWriter, r *http.Request, c *CustomerRepository) {
	// new Entry var
	var newCustomer Customer
	// read the request body & parse JSON body
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // early  fail to avoid duplicate writing of status code
	}

	// insert new Entry into Customers map
	c.Save(newCustomer)

	// get new entry from database for validation pupose
	customer, err := c.FindByID(newCustomer.ID)
	if err != nil {
		errors.New("could retrieve newly added item from database" + newCustomer.ID.String())
		w.WriteHeader(http.StatusNotFound)
	}
	// output new entry
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// Retrieve single customer by ID in /customer/{id}
func getCustomer(w http.ResponseWriter, r *http.Request, c *CustomerRepository) {
	vars := mux.Vars(r)
	CustomerID, _ := uuid.Parse(vars["id"])
	foundCustomer, err := c.FindByID(CustomerID)

	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundCustomer)
}

// update Customer by ID
func updateCustomer(w http.ResponseWriter, r *http.Request, c *CustomerRepository) {
	vars := mux.Vars(r)
	CustomerID, _ := uuid.Parse(vars["id"])

	// new Entry var
	var updateCustomer Customer
	// read the request body & parse JSON body
	err := json.NewDecoder(r.Body).Decode(&updateCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // early  fail to avoid duplicate writing of status code
	}
	updateCustomer.ID = CustomerID
	// insert new Entry into Customers map

	updatedCustomer, err := c.Save(updateCustomer)
	if err != nil {
		http.Error(w, "Customer not found: "+updatedCustomer.ID.String(), http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// remove a single Customer by ID
func deleteCustomer(w http.ResponseWriter, r *http.Request, c *CustomerRepository) {
	// Get ID from Var ULR
	vars := mux.Vars(r)
	customerID, _ := uuid.Parse(vars["id"])

	err := c.DeleteByID(customerID)
	// if not found fail else remove
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Customer %v was removed", customerID)
	}
}

// better readability of wrapper function
type CustomHandler func(http.ResponseWriter, *http.Request, *CustomerRepository)

// wrap HandleFunc to handover CustomerRepository Pointer into called function like getCustomers,...
func wrap(f CustomHandler, c *CustomerRepository) func(http.ResponseWriter, *http.Request) {
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
	router.HandleFunc("/customers", wrap(getCustomers, customerRepository)).Methods("GET")
	router.HandleFunc("/customers", wrap(addCustomer, customerRepository)).Methods("POST")
	router.HandleFunc("/customers/{id}", wrap(getCustomer, customerRepository)).Methods("GET")
	router.HandleFunc("/customers/{id}", wrap(updateCustomer, customerRepository)).Methods("PUT")
	router.HandleFunc("/customers/{id}", wrap(deleteCustomer, customerRepository)).Methods("DELETE")

	fmt.Println("Starting the server on: ", url)
	err := http.ListenAndServe(url, router)
	if err != nil {
		panic(err)
	}
}
