package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/cecepsprd/go-api-native/model"
	"github.com/cecepsprd/go-api-native/service"
)

type ProductHandler struct {
	DB *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{db}
}

func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Query().Get("id") != "" {
			p.GetProductByID(w, r)
			return
		}
		p.GetProducts(w, r)
		return
	}

	if r.Method == "POST" {
		p.AddProduct(w, r)
		return
	}
}

func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	products, err := service.GetProductsService(ctx, p.DB)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (p *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	product, err := service.GetProduct(ctx, p.DB, int64(id))
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (p *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var pr model.Product
	_ = json.Unmarshal(body, &pr)

	product, err := service.AddProductService(ctx, p.DB, pr)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
