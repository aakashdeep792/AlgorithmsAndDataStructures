package main

import (
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
	
	
	r := mux.NewRouter().PathPrefix("aakash")

	r.Handle().Methods(http.MethodGet).

	s:= http.Server{}


	

	mux.Handle("/", http.HandlerFunc(dummy)).

}
