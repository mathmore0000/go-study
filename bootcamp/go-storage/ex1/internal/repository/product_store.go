package repository

import (
	"app/internal"
	"database/sql"
	"errors"
)

// NewRepositoryProductStore creates a new repository for products.
func NewRepositoryProductStore(db *sql.DB) (r *RepositoryProductStore) {
	r = &RepositoryProductStore{
		db: db,
	}
	return
}

// RepositoryProductStore is a repository for products.
type RepositoryProductStore struct {
	// st is the underlying store.
	st internal.StoreProduct
	db *sql.DB
}

// FindById finds a product by id.
func (r *RepositoryProductStore) FindById(id int) (p internal.Product, err error) {
	// read all products
	row := r.db.QueryRow("SELECT p.id, p.name, p.quantity, p.code_value, p.is_published, p.expiration, p.price FROM products p WHERE id = ?", id)
	if row.Err() != nil {
		return internal.Product{}, err
	}

	// read product
	err = row.Scan(&p.Id, &p.ProductAttributes.Name, &p.ProductAttributes.Quantity, &p.ProductAttributes.CodeValue, &p.ProductAttributes.IsPublished, &p.ProductAttributes.Expiration, &p.ProductAttributes.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.Product{}, internal.ErrRepositoryProductNotFound
		}

		return internal.Product{}, err
	}

	return
}

// Save saves a product.
func (r *RepositoryProductStore) Save(p *internal.Product) (err error) {
	result, err := r.db.Exec("INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?,?,?,?,?,?)",
		p.ProductAttributes.Name, p.ProductAttributes.Quantity, p.ProductAttributes.CodeValue, p.ProductAttributes.IsPublished, p.ProductAttributes.Expiration, p.ProductAttributes.Price)
	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	p.Id = int(id)

	return
}

// UpdateOrSave updates or saves a product.
func (r *RepositoryProductStore) UpdateOrSave(p *internal.Product) (err error) {
	err = r.Update(p)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryProductNotFound) {
			return r.Save(p)
		}
		return
	}
	return
}

// Update updates a product.
func (r *RepositoryProductStore) Update(p *internal.Product) (err error) {
	if _, err = r.FindById(p.Id); err != nil {
		return
	}
	_, err = r.db.Exec("UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?",
		p.ProductAttributes.Name, p.ProductAttributes.Quantity, p.ProductAttributes.CodeValue, p.ProductAttributes.IsPublished, p.ProductAttributes.Expiration, p.ProductAttributes.Price, p.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.ErrRepositoryProductNotFound
		}
		return
	}
	return nil
}

// Delete deletes a product.
func (r *RepositoryProductStore) Delete(id int) (err error) {
	if _, err = r.FindById(id); err != nil {
		return
	}

	_, err = r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return internal.ErrRepositoryProductNotFound
		}
		return
	}
	return
}
