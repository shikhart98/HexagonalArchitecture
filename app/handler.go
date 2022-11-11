package app

import (
	"HexagonalArch/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type CustomerHandlers struct {
	customerService service.CustomerService
}

func (h *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer {
	//	{"Shikhar", "New Delhi", "110088"},
	//	{"Amit", "Lucknow", "229900"},
	//}
	customers, _ := h.customerService.GetAllCustomers()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := h.customerService.GetCustomer(customerId)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}