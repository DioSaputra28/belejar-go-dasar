package service

import (
	"github.com/DioSaputra28/belejar-go-dasar/internal/category/model"
	"github.com/DioSaputra28/belejar-go-dasar/internal/category/repository"
)

type CategoryService interface {
	GetAll() ([]model.Category, error)
	CreateCategory(category *model.Category) error
	GetCategoryById(id int) (model.Category, error)
	UpdateCategory(category *model.Category) error
	DeleteCategory(id int) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryService) CreateCategory(category *model.Category) error {
	return s.repo.CreateCategory(category)
}

func (s *categoryService) GetCategoryById(id int) (model.Category, error) {
	return s.repo.GetCategoryById(id)
}

func (s *categoryService) UpdateCategory(category *model.Category) error {
	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id)
}
