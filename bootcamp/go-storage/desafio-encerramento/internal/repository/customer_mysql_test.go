package repository

import (
	"app/internal"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	txdb.Register("txdb", "mysql", "root:root@tcp(localhost:3306)/fantasy_products?parseTime=true")
}

func TestCustomersMySQL_GetTop5SpentMost(t *testing.T) {
	t.Log("Starting TestCustomersMySQL_GetTop5SpentMost")

	db, err := sql.Open("txdb", "conn-get-top5-customers")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	t.Log("Database connection opened")

	t.Log("Test data seeded")

	rpCustomer := NewCustomersMySQL(db)
	t.Log("Repository initialized")

	c, err := rpCustomer.GetTop5SpentMost()
	t.Log("Called GetTop5SpentMost")
	if err != nil {
		t.Fatalf("Error in GetTop5SpentMost: %v", err)
	}
	t.Log("GetTop5SpentMost completed")

	expected := []internal.CustomerMostSpent{
		{FirstName: "Lannie", LastName: "Tortis", Ammount: 58513.55},
		{FirstName: "Jasen", LastName: "Crowcum", Ammount: 48291.03},
		{FirstName: "Elvina", LastName: "Ovell", Ammount: 43590.75},
		{FirstName: "Lazaro", LastName: "Anstis", Ammount: 40792.06},
		{FirstName: "Wilden", LastName: "Oaten", Ammount: 39786.79},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, c)
	t.Log("TestCustomersMySQL_GetTop5SpentMost completed successfully")
}

func TestCustomersMySQL_GetTotalGroupedByCondition(t *testing.T) {
	t.Log("Starting TestCustomersMySQL_GetTotalGroupedByCondition")

	db, err := sql.Open("txdb", "conn-get-total-grouped")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	t.Log("Database connection opened")

	t.Log("Test data seeded")

	rpCustomer := NewCustomersMySQL(db)
	t.Log("Repository initialized")

	c, err := rpCustomer.GetTotalGroupedByCondition()
	t.Log("Called GetTotalGroupedByCondition")
	if err != nil {
		t.Fatalf("Error in GetTop5SpentMost: %v", err)
	}
	t.Log("GetTotalGroupedByCondition completed")

	expected := map[int]float32{
		0: 605929.10,
		1: 716792.33,
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, c)
	t.Log("TestCustomersMySQL_GetTotalGroupedByCondition completed successfully")
}
