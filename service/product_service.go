package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/cecepsprd/go-api-native/model"
)

func GetProductsService(c context.Context, db *sql.DB) ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	products, err := model.GetProducts(db, ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProduct(c context.Context, db *sql.DB, id int64) (p model.Product, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	p, err = model.GetProductByID(db, ctx, id)
	if err != nil {
		return
	}

	return
}

func AddProductService(c context.Context, db *sql.DB, p model.Product) (lastID int64, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	lastID, err = model.AddProduct(db, ctx, p)
	if err != nil {
		return 0, err
	}

	return
}
