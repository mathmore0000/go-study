package repository

import (
	"database/sql"

	"app/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// GetTop5SpentMost returns the top 5 customers that spent the most
func (r *CustomersMySQL) GetTop5SpentMost() (c []internal.CustomerMostSpent, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT c.first_name, c.last_name, ROUND(SUM(i.total), 2) FROM customers c INNER JOIN invoices i ON c.id = i.customer_id GROUP BY c.id ORDER BY SUM(i.total) DESC LIMIT 5;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.CustomerMostSpent
		// scan the row into the customer
		err := rows.Scan(&cs.FirstName, &cs.LastName, &cs.Ammount)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *CustomersMySQL) GetTotalGroupedByCondition() (t map[int]float32, err error) {
	t = make(map[int]float32)
	// execute the query
	rows, err := r.db.Query("SELECT c.condition, ROUND(SUM(i.total),2) FROM customers c INNER JOIN invoices i ON c.id = i.customer_id GROUP BY `condition`;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var condition int
		var total float32
		// scan the row into the customer
		err := rows.Scan(&condition, &total)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		t[condition] = total
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*c).Id = int(id)

	return
}
