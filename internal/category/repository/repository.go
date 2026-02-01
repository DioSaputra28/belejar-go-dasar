package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/DioSaputra28/belejar-go-dasar/internal/category/model"
)

type CategoryRepository interface {
	GetAll() ([]model.Category, error)
	CreateCategory(category *model.Category) error
	GetCategoryById(id int) (model.Category, error)
	UpdateCategory(category *model.Category) error
	DeleteCategory(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll() ([]model.Category, error) {
	rows, err := r.db.Query("SELECT id, name, description, created_at FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *categoryRepository) CreateCategory(category *model.Category) error {
	query := `INSERT INTO category (name, description, created_at) VALUES ($1, $2, $3) RETURNING id`

	category.CreatedAt = time.Now()

	err := r.db.QueryRow(query, category.Name, category.Description, category.CreatedAt).Scan(&category.ID)
	return err
}

func (r *categoryRepository) GetCategoryById(id int) (model.Category, error) {
	var category model.Category
	query := `SELECT id, name, description, created_at FROM category WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt)
	if err == sql.ErrNoRows {
		return model.Category{}, errors.New("category not found")
	}
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (r *categoryRepository) UpdateCategory(category *model.Category) error {
	query := `UPDATE category SET name = $1, description = $2 WHERE id = $3`

	result, err := r.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("category not found")
	}

	return nil
}

func (r *categoryRepository) DeleteCategory(id int) error {
	query := `DELETE FROM category WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("category not found")
	}

	return nil
}
