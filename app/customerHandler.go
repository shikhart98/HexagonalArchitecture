package app

import (
	"HexagonalArch/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


type CustomerHandlers struct {
	customerService service.CustomerService
}

func (h *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status_filter := r.URL.Query().Get("status")
	fmt.Println("Status filter: " + status_filter)
	filters := map[string] string {
		"status": status_filter,
	}
	customers, err := h.customerService.GetAllCustomers(&filters)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (h *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := h.customerService.GetCustomer(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse (w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}