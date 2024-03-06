package store

import (
	"algoDS/goSpecificQuestions/crudOps/model"
	"context"
)

type Storage interface {
	GetProductFromDB(ctx context.Context) ([]model.ProductStocks, error)
	RegisterProductDetailsInDB(ctx context.Context, p model.Product)
	UpdateProductDetailsInDB(ctx context.Context, p model.Product)
	DeleteProductInDB(ctx context.Context, p model.Product)
	RegisterLocationInDB(ctx context.Context, l model.Location)
	UpdateLocationInDB(ctx context.Context, l model.Location)
	DeleteLocationInDB(ctx context.Context, l model.Location)
	IncrProductStockInDB(ctx context.Context)
	DecrProductStockInDB(ctx context.Context)
}
