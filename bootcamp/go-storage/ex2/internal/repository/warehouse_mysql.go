package repository

import (
	"app/internal"
	"database/sql"
	"fmt"
)

// NewWarehousesMySQL returns a new instance of WarehousesMySQL
func NewWarehousesMySQL(db *sql.DB) *WarehousesMySQL {
	return &WarehousesMySQL{
		db: db,
	}
}

// WarehousesMySQL is a struct that represents a warehouse repository
type WarehousesMySQL struct {
	// db is the database connection
	db *sql.DB
}

// GetAll returns all warehouses
func (r *WarehousesMySQL) GetAll() (w []internal.Warehouse, err error) {
	// execute the query
	rows, err := r.db.Query(
		"SELECT `id`, `name`, `adress`" +
			"FROM `warehouses`",
	)
	if err != nil {
		return
	}

	// scan the rows into the warehouses
	for rows.Next() {
		var wh internal.Warehouse
		err = rows.Scan(&wh.ID, &wh.Name, &wh.Address)
		if err != nil {
			return
		}
		w = append(w, wh)
	}

	return
}

func (r *WarehousesMySQL) GetProductsCount(id *int) (w []internal.WarehouseProductCount, err error) {
	var rows *sql.Rows
	if id != nil {
		// execute the query
		rows, err = r.db.Query(
			"SELECT warehouses.name, COUNT(products.id) "+
				"FROM warehouses "+
				"LEFT JOIN products ON warehouses.id = products.id_warehouse "+
				"where warehouses.id = ? "+
				"GROUP BY warehouses.id;",
			id,
		)
	} else {
		rows, err = r.db.Query(
			"SELECT warehouses.name, COUNT(products.id) " +
				"FROM warehouses " +
				"LEFT JOIN products ON warehouses.id = products.id_warehouse " +
				"GROUP BY warehouses.id;")
	}
	if err != nil {
		fmt.Println("repository/warehouse_mysql.go: ", err)
		return
	}

	// scan the row into the warehouse
	for rows.Next() {
		var wp internal.WarehouseProductCount
		err = rows.Scan(&wp.Name, &wp.ProductCount)
		if err != nil {
			return
		}
		w = append(w, wp)
	}
	return
}

// GetOne returns a warehouse by id
func (r *WarehousesMySQL) GetOne(id int) (p internal.Warehouse, err error) {
	// execute the query
	row := r.db.QueryRow(
		"SELECT w.id, w.name, w.adress, w.telephone, w.capacity FROM warehouses w WHERE ID = ?",
		id,
	)
	if err = row.Err(); err != nil {
		return
	}

	// scan the row into the warehouse
	err = row.Scan(&p.ID, &p.Name, &p.Address, &p.Telephone, &p.Capacity)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrWarehouseNotFound
		}
		return
	}

	return
}

// Store stores a warehouse
func (r *WarehousesMySQL) Store(p *internal.Warehouse) (err error) {
	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `warehouses` (`name`, `adress`, `telephone`, `capacity`) "+
			"VALUES (?, ?, ?, ?)",
		p.Name, p.Address, p.Telephone, p.Capacity,
	)
	if err != nil {
		return
	}

	// get the last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	p.ID = int(id)

	return
}
