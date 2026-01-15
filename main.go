package main

import (
	"database/sql"

	db2 "github.com/mauriciovictor/curso-hexagonal/adapters/db"
	"github.com/mauriciovictor/curso-hexagonal/application"
)

func main() {
	DB, _ := sql.Open("sqlite3", "sqlite.db")
	productDBAdapter := db2.NewProductDB(DB)
	productService := application.NewProductService(productDBAdapter)

	product, _ := productService.Save("Product Example", 100)
	_, err := productService.Enable(product)

	if err != nil {
		return
	}
}
