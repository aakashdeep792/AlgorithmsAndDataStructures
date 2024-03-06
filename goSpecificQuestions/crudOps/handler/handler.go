package handler

import (
	"algoDS/goSpecificQuestions/crudOps/model"
	"algoDS/goSpecificQuestions/crudOps/store"
	"encoding/json"
	"net/http"
)

type Inventory struct {
	s store.Storage
}

func NewInventoryHandler(s store.Storage) *Inventory {
	return &Inventory{s: s}
}

func (i *Inventory) GetProducts(w http.ResponseWriter, r *http.Request) {
	val, err := i.s.GetProductFromDB(r.Context())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	bytes, _ := json.MarshalIndent(val, "", "  ")
	w.Write(bytes)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

}

func RegisterProductDetails(p *model.Product) interface{} {
	return nil

}

func UpdateProductDetails() {

}

func DeleteProduct() {

}

func RegisterLocation() { // can be seller

}

func UpdateLocation() {

}

func DeleteLocation() {

}

func IncrProductStock() {

}

func DecrProductStock() {

}
