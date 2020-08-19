package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Product struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Stock     int32  `json:"stock"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Products []Product

func GetProducts(db *sql.DB, c context.Context) (products Products, err error) {
	query := "SELECT id, name, price, stock, created_at, updated_at FROM product limit 10"

	rows, err := db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product

		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			fmt.Println("Model 2")
			log.Print(err)
			return
		}

		products = append(products, p)
	}

	return products, nil
}

func GetProductByID(db *sql.DB, c context.Context, id int64) (p Product, err error) {
	query := "SELECT id, name, price, stock, created_at, updated_at FROM product where id = ?"
	err = db.QueryRowContext(c, query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)

	return
}

func AddProduct(db *sql.DB, c context.Context, p Product) (int64, error) {
	query := "INSERT INTO product(name, price, stock, created_at, updated_at) VALUES(?,?,?,?,?)"

	stmt, err := db.PrepareContext(c, query)
	if err != nil {
		fmt.Println("1")
		return 0, err
	}

	res, err := stmt.ExecContext(c, p.Name, p.Price, p.Stock, time.Now(), time.Now())
	if err != nil {
		fmt.Println("2")
		return 0, err
	}

	lastID, _ := res.LastInsertId()
	return lastID, nil
}
