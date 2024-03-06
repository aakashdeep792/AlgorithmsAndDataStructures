package model

type Location struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"Address"`
	Pin     int    `json:"Pin"`
}

type ProductStocks struct {
	Product
	Locs []LocationProductCount `json:"locations"`
}

type LocationProductCount struct {
	Location
	Count int `json:"count"`
}
