package app

import (
	"banking/domain"
	"banking/helper"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	// wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	hostname := helper.GetEnvVariable("SERVER_HOST")
	port := helper.GetEnvVariable("SERVER_PORT")

	// starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", hostname, port), router))
}