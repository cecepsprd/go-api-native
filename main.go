package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cecepsprd/go-api-native/config"
	"github.com/cecepsprd/go-api-native/handler"
)

func main() {

	//connect to db
	db, err := config.DBConnect()
	if err != nil {
		log.Fatalf("Failed to connect database (Error: %s)\n", err)
	}
	defer db.Close()

	//Handler
	productHandler := handler.NewProductHandler(db)
	http.HandleFunc("/product", productHandler.ServeHTTP)

	//Starting server
	log.Println("Starting server on port 9090")
	if err := http.ListenAndServe(":9292", nil); err == nil {
		log.Printf("Error Starting server on port : %s \n", ":9292")
		os.Exit(1)
	}

}
