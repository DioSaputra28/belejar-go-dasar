package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/DioSaputra28/belejar-go-dasar/internal/produk/model"
)

type ProductRepository interface {
	GetAll(name string) ([]model.Product, error)
	CreateProduct(product *model.Product) error
	GetProductById(id int) (model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(name string) ([]model.Product, error) {
	var rows *sql.Rows
	var err error

	if name == "" {
		rows, err = r.db.Query("SELECT id, name, price, stock, created_at FROM product")
	} else {
		rows, err = r.db.Query("SELECT id, name, price, stock, created_at FROM product WHERE name ILIKE $1", "%"+name+"%")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) CreateProduct(product *model.Product) error {
	query := `INSERT INTO product (name, price, stock, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	product.CreatedAt = time.Now()

	err := r.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CreatedAt).Scan(&product.ID)
	return err
}

func (r *productRepository) GetProductById(id int) (model.Product, error) {
	var product model.Product
	query := `SELECT id, name, price, stock, created_at FROM product WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt)
	if err == sql.ErrNoRows {
		return model.Product{}, errors.New("product not found")
	}
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (r *productRepository) UpdateProduct(product *model.Product) error {
	query := `UPDATE product SET name = $1, price = $2, stock = $3 WHERE id = $4`

	result, err := r.db.Exec(query, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *productRepository) DeleteProduct(id int) error {
	query := `DELETE FROM product WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}
