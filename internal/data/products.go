package data

import (
	"database/sql"
	"errors"
	"fmt"
)

type ProductsModel struct {
	DB *sql.DB
}

func (p ProductsModel) Insert(product *Product) error {
	query := `
	INSERT INTO products (product_id, product_name) 
	VALUES ($1,$2)`

	args := []interface{}{product.ID, product.Name}

	_, err := p.DB.Exec(query, args...)
	if err != nil {
		err = fmt.Errorf("error inserting product in db: %w", err)
	}
	return err
}

func (p ProductsModel) Exists(productID string) (bool, error) {
	query := `
	SELECT product_name 
	FROM products
	WHERE product_id = $1`

	var name string
	err := p.DB.QueryRow(query, productID).Scan(&name)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return false, ErrRecordNotFound
		default:
			return false, err
		}
	}

	return true, nil
}
