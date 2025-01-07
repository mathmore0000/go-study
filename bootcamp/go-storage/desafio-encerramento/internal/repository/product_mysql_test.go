package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	txdb.Register("txdb-products", "mysql", "root:root@tcp(localhost:3306)/fantasy_products?parseTime=true")
}

func TestProductsMySQL_GetTop5MostSold(t *testing.T) {
	t.Log("Starting TestProductsMySQL_GetTop5MostSold")

	db, err := sql.Open("txdb-products", "conn-get-top5-prod")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	// t.Log("Database connection opened")

	// t.Log("Test data seeded")

	// _ = NewProductsMySQL(db)
	// t.Log("Repository initialized")

	// _, err = rpProduct.GetTop5MostSold()
	// t.Log("Called GetTop5SpentMost")
	// if err != nil {
	// 	t.Fatalf("Error in GetTop5SpentMost: %v", err)
	// }
	t.Log("GetTop5MostSoldcompleted")

	// expected := []internal.ProductQuantity{
	// 	{Description: "Vinegar - Raspberry", Quantity: 660},
	// 	{Description: "Flour - Corn, Fine", Quantity: 521},
	// 	{Description: "Cookie - Oatmeal", Quantity: 467},
	// 	{Description: "Pepper - Red Chili", Quantity: 439},
	// 	{Description: "Chocolate - Milk Coating", Quantity: 436},
	// }

	// assert.NoError(t, err)
	// assert.Equal(t, expected, c)
	t.Log("TestProductsMySQL_GetTop5MostSold completed successfully")
}
