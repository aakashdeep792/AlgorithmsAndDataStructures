package handler

import "algoDS/goSpecificQuestions/crudOps/model"

type ProductStocks struct {
	model.Product
	Locs []LocationProductCount `json:"locations"`
}

type LocationProductCount struct {
	model.Location
	Count int `json:"count"`
}
