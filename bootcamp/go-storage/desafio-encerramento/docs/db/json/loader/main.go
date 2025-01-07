package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// set up mysql
// go get github.com/go-sql-driver/mysql
// mysql -u root -p -P 3306 -h localhost

// run customers.json to mysql
// run invoices.json to mysql
// run products.json to mysql
// run sales.json to mysql

type Customer struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

type Invoice struct {
	Id         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	CustomerId int     `json:"customer_id"`
	total      float32 `json:"total"`
}

type Product struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type Sale struct {
	Id        int `json:"id"`
	InvoiceId int `json:"invoice_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func main() {
	storage, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/fantasy_products?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer storage.Close()
	if err := storage.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Conectado ao mysql")

	customers, err := loadCustomers("../customers.json")
	if err != nil {
		panic(err)
	}

	invoices, err := loadInvoices("../invoices.json")
	if err != nil {
		panic(err)
	}

	products, err := loadProducts("../products.json")
	if err != nil {
		panic(err)
	}

	sales, err := loadSales("../sales.json")
	if err != nil {
		panic(err)
	}

	err = saveCustomers(storage, customers)
	if err != nil {
		panic(err)
	}
	err = saveProducts(storage, products)
	if err != nil {
		panic(err)
	}
	err = saveInvoices(storage, invoices)
	if err != nil {
		panic(err)
	}
	err = saveSales(storage, sales)
	if err != nil {
		panic(err)
	}
}

func saveCustomers(storage *sql.DB, customers []Customer) error {
	sql := "INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES"

	for _, customer := range customers {
		sql += fmt.Sprintf(` (%d, "%s", "%s", %d),`, customer.Id, customer.FirstName, customer.LastName, customer.Condition)
	}

	sql = sql[:len(sql)-1]

	_, err := storage.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func saveInvoices(storage *sql.DB, invoices []Invoice) error {
	sql := "INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES"

	for _, invoice := range invoices {
		sql += fmt.Sprintf(` (%d, "%s", %d, %f),`, invoice.Id, invoice.Datetime, invoice.CustomerId, invoice.total)
	}

	sql = sql[:len(sql)-1]

	_, err := storage.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func saveProducts(storage *sql.DB, products []Product) error {
	sql := "INSERT INTO products (`id`, `description`, `price`) VALUES"

	for _, product := range products {
		sql += fmt.Sprintf(` (%d, "%s", %f),`, product.Id, product.Description, product.Price)
	}

	sql = sql[:len(sql)-1]

	_, err := storage.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func saveSales(storage *sql.DB, sales []Sale) error {
	sql := "INSERT INTO sales (`id`, `invoice_id`, `product_id`, `quantity`) VALUES"

	for _, sale := range sales {
		sql += fmt.Sprintf(` (%d, %d, %d, %d),`, sale.Id, sale.InvoiceId, sale.ProductId, sale.Quantity)
	}

	sql = sql[:len(sql)-1]

	_, err := storage.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func loadCustomers(path string) ([]Customer, error) {
	customers := make([]Customer, 0)

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&customers)
	if err != nil {
		return nil, err
	}

	return customers, err
}

func loadInvoices(path string) ([]Invoice, error) {
	invoices := make([]Invoice, 0)

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&invoices)
	if err != nil {
		return nil, err
	}
	return invoices, err
}

func loadProducts(path string) ([]Product, error) {
	products := make([]Product, 0)

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&products)
	if err != nil {
		return nil, err
	}
	return products, err
}

func loadSales(path string) ([]Sale, error) {
	sales := make([]Sale, 0)

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&sales)
	if err != nil {
		return nil, err
	}
	return sales, err
}
