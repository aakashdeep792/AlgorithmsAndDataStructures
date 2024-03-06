package store

import (
	"algoDS/goSpecificQuestions/crudOps/model"
	"context"
	"sync"
	"time"
)

type inMemoryStore struct {
	prods sync.Map // store Products

	locations map[int][]model.LocationProductCount
	//productID ->[]LocationProductCount
}

func NewInMemoryStore() Storage {
	tmp := &inMemoryStore{
		prods:     sync.Map{},
		locations: make(map[int][]model.LocationProductCount),
	}

	p1 := model.Product{ID: 1, Name: "iphone15", Decription: "cpu:M1,ram:8GB,storage:128GB", CreatedAt: time.Now(), Size: model.SMALL}
	p2 := model.Product{ID: 2, Name: "ROG", Decription: "cpu:i78358H,ram:8GB,storage:128GB", CreatedAt: time.Now(), Size: model.MEDIUM}
	l1 := model.LocationProductCount{Location: model.Location{ID: 1, Name: "Reliance Digital Store", Address: "Bangalore", Pin: 893830}, Count: 10}
	l2 := model.LocationProductCount{Location: model.Location{ID: 2, Name: "Croma Store", Address: "Mumbai", Pin: 726930}, Count: 34}

	tmp.prods.Store(p1.ID, p1)
	tmp.prods.Store(p2.ID, p2)
	tmp.locations[p1.ID] = []model.LocationProductCount{l1}
	tmp.locations[p2.ID] = []model.LocationProductCount{l1, l2}

	return tmp
}

func (m *inMemoryStore) GetProductFromDB(ctx context.Context) ([]model.ProductStocks, error) {
	val := make([]model.ProductStocks, 0)

	m.prods.Range(func(k, v any) bool {
		tmp, ok := v.(model.Product)
		if !ok {
			return true
		}

		val = append(val, model.ProductStocks{Product: tmp})
		return true
	})

	for i := 0; i < len(val); i++ {
		val[i].Locs = m.locations[val[i].Product.ID]
	}

	return val, nil
}

func (m *inMemoryStore) RegisterProductDetailsInDB(ctx context.Context, p model.Product) {}
func (m *inMemoryStore) UpdateProductDetailsInDB(ctx context.Context, p model.Product)   {}
func (m *inMemoryStore) DeleteProductInDB(ctx context.Context, p model.Product)          {}
func (m *inMemoryStore) RegisterLocationInDB(ctx context.Context, l model.Location)      {}
func (m *inMemoryStore) UpdateLocationInDB(ctx context.Context, l model.Location)        {}
func (m *inMemoryStore) DeleteLocationInDB(ctx context.Context, l model.Location)        {}
func (m *inMemoryStore) IncrProductStockInDB(ctx context.Context)                        {}
func (m *inMemoryStore) DecrProductStockInDB(ctx context.Context)                        {}
