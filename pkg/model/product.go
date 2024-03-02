package model

import "crud_product/pkg/dto"

type Product struct {
	Id          int     `json:"id"`
	ImageUrl    string  `json:"image_url"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductRepository interface {
	GetAllProduct() ([]Product, error)
	CreateProduct(req Product) error
	UpdateProduct(id int, req Product) error
	DeleteProduct(id int) error
	GetProductById(id int) (Product, error)
}
type ProductUsecase interface {
	GetAllProduct() ([]Product, error)
	GetProductById(id int) (Product, error)
	CreateProduct(req dto.ProductRequest) error
	UpdateProduct(id int, req dto.ProductRequest) error
	DeleteProduct(id int) error
}
