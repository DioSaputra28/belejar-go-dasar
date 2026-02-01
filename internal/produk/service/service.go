package service

import (
	"github.com/DioSaputra28/belejar-go-dasar/internal/produk/model"
	"github.com/DioSaputra28/belejar-go-dasar/internal/produk/repository"
)

type ProductService interface {
	GetAll() ([]model.Product, error)
	CreateProduct(product *model.Product) error
	GetProductById(id int) (model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) CreateProduct(product *model.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) GetProductById(id int) (model.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *productService) UpdateProduct(product *model.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
