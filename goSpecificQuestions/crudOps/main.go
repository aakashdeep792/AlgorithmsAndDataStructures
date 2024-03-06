package main

import (
	"algoDS/goSpecificQuestions/crudOps/handler"
	"algoDS/goSpecificQuestions/crudOps/store"
	"net/http"

	"github.com/gorilla/mux"
)

func dummy(w http.ResponseWriter, r *http.Request) {

}

// Inventory Management

// grpc
// migration
// db orm
// pagination
// crud
//caching   -> memcache and redis
// data compression
// rate limiting

func main() {
	r := mux.NewRouter()
	// r.Handler(i.GetProducts())

	i := handler.NewInventoryHandler(store.NewInMemoryStore())
	r.HandleFunc("/listProducts", i.GetProducts).Methods(http.MethodGet)

	// http.Handle(/,r)

	svr := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	svr.ListenAndServe()

}
