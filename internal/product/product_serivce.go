package product

import (
	"errors"
)

type ProductDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type ProductService struct {
	// Sau này bạn bơm Repository hoặc Supabase Client vào đây
}

func CreateProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) FindProductById(id string) (*ProductDTO, error) {
	if id == "" {
		return nil, errors.New("ID can not be blank")
	}
	if id == "404" {
		return nil, errors.New("Can not find product")
	}
	product := &ProductDTO{
		ID:          id,
		Name:        "Laptop",
		Description: "This is a laptop",
		Category:    "Electronics",
	}

	return product, nil
}
