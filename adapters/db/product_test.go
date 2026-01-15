package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/mauriciovictor/curso-hexagonal/adapters/db"
	"github.com/mauriciovictor/curso-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string PRIMARY KEY, "name" string, "price" float, "status" string)`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("abc", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDB := db.NewProductDB(Db)
	product, err := productDB.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDB(Db)

	product := application.NewProduct("Product Test 2", 1)
	product.Status = application.ENABLED

	result, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, 1.0, result.GetPrice())
	require.Equal(t, "Product Test 2", result.GetName())
	require.Equal(t, "enabled", result.GetStatus())

	product.Status = application.DISABLED
	result, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, "disabled", result.GetStatus())
}
