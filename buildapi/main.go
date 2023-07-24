package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Modle structure for Customer
type Customer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Car  *Car   `json:"car"`
}

// Modle structure for Car
type Car struct {
	BrandName string `json:"brand_name"`
	Color     string `json:"color"`
}

// fake DB
var customers []Customer

// middleware, helper - file
func (c *Customer) IsEmpty() bool {
	// return c.Name == "" && c.Age == 0
	return c.Name == "" && c.Age == 0
}

func main() {
}

// servce home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello this is home page</h1>"))
}

// get All Customer
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all customers")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// get Customer by Name
func getOneCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Customer")
	w.Header().Set("Content-Type", "application/json")

	// grep one customer
	params := mux.Vars(r)

	// loop through customes, find matching id and return the response
	for _, customer := range customers {
		if customer.Name == params["name"] {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	json.NewEncoder(w).Encode("No Customer Found")
	return
}

// Create Customer
func createOneCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Customer")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is missing")
	}

	var customer Customer
	_ = json.NewDecoder(r.Body).Decode(&customer)
	if customer.IsEmpty() {
		json.NewEncoder(w).Encode("No data found inside JSON")
		return
	}

	//generate name
	//append customer into customers
	rand.Seed(time.Now().UnixNano())
	customer.Name = "Nan Oo Hlaing"
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer)
	return
}

// Update Customer
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Customer")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	//loop
	for index, customer := range customers {
		if customer.Age == params["age"] {
			customers = append(customers[:index], customers[index+1:]...)
			var customer Customer
			_ = json.NewDecoder(r.Body).Decode(&customer)
			customer.Age = params["id"]
			customers = append(customers, customer)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
}

// Delete Customer
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delte One Customer")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove
	for index, customer := range customers {
		if customer.Name == params["name"] {
			customers = append(customers[:index], customers[index+1]...)
			break
		}
	}
}
