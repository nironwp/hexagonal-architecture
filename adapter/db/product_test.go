package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/nironwp/hexagonal-architecture/adapter/db"
	"github.com/nironwp/hexagonal-architecture/application"
	"github.com/stretchr/testify/assert"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string,
		"quantity" int
	)
	`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Candy", 1, "disabled", 1)`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestProductDb_Get(t *testing.T) {
	setUp()

	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")

	assert.Nil(t, err)

	assert.Equal(t, "abc", product.GetID())
	assert.Equal(t, "Candy", product.GetName())
	assert.Equal(t, 1.0, product.GetPrice())
	assert.Equal(t, 1, product.GetQuantity())
}

func TestProductDb_Save_Create(t *testing.T) {
	setUp()

	defer Db.Close()
	productDb := db.NewProductDb(Db)

	inputDto := application.NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1,
	}

	product, err := application.NewProduct(inputDto)

	assert.Nil(t, err)

	_, err = productDb.Save(product)

	assert.Nil(t, err)

	productGet, err := productDb.Get(product.GetID())

	assert.Equal(t, productGet.GetID(), product.GetID())
	assert.Equal(t, productGet.GetName(), product.GetName())
	assert.Equal(t, productGet.GetPrice(), product.GetPrice())
	assert.Equal(t, productGet.GetQuantity(), product.GetQuantity())
	assert.Equal(t, productGet.GetStatus(), product.GetStatus())
}

func TestProductDb_Save_Update(t *testing.T) {
	setUp()

	defer Db.Close()
	productDb := db.NewProductDb(Db)

	inputDto := application.NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1,
	}

	product, err := application.NewProduct(inputDto)

	assert.Nil(t, err)

	_, err = productDb.Save(product)

	assert.Nil(t, err)

	product.ChangePrice(2)

	_, err = productDb.Save(product)

	assert.Nil(t, err)
	productGet, err := productDb.Get(product.GetID())

	assert.Equal(t, productGet.GetID(), product.GetID())
	assert.Equal(t, productGet.GetName(), product.GetName())
	assert.Equal(t, productGet.GetPrice(), product.GetPrice())
	assert.Equal(t, productGet.GetQuantity(), product.GetQuantity())
	assert.Equal(t, productGet.GetStatus(), product.GetStatus())

}
