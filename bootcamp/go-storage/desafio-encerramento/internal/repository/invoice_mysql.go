package repository

import (
	"database/sql"

	"app/internal"
)

// NewInvoicesMySQL creates new mysql repository for invoice entity.
func NewInvoicesMySQL(db *sql.DB) *InvoicesMySQL {
	return &InvoicesMySQL{db}
}

// InvoicesMySQL is the MySQL repository implementation for invoice entity.
type InvoicesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// UpdateTotal updates the total of all invoices.
func (r *InvoicesMySQL) UpdateTotal() (err error) {
	// execute the query
	_, err = r.db.Exec(
		"update invoices i JOIN " +
			"(SELECT " +
			"i.id AS invoice_id, " +
			"SUM(ROUND(p.price * s.quantity, 2)) AS total_invoice " +
			"FROM " +
			"sales s " +
			"INNER JOIN " +
			"invoices i ON s.invoice_id = i.id " +
			"INNER JOIN " +
			"products p ON s.product_id = p.id " +
			"GROUP BY " +
			"i.id) calc_total ON i.id = calc_total.invoice_id " +
			"SET total = calc_total.total_invoice;")
	return
}

// FindAll returns all invoices from the database.
func (r *InvoicesMySQL) FindAll() (i []internal.Invoice, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var iv internal.Invoice
		// scan the row into the invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		// append the invoice to the slice
		i = append(i, iv)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the invoice into the database.
func (r *InvoicesMySQL) Save(i *internal.Invoice) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
		(*i).Datetime, (*i).Total, (*i).CustomerId,
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
	(*i).Id = int(id)

	return
}
