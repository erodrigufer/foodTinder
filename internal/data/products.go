package data

import (
	"database/sql"
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
