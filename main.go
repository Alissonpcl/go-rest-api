package main

import (
	"github.com/alissonpcl/go-rest-api/database"
	"github.com/alissonpcl/go-rest-api/product"
	"github.com/alissonpcl/go-rest-api/receipt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	product.SetupRoutes(apiBasePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
