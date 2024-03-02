package usecase

import (
	"crud_product/pkg/dto"
	"crud_product/pkg/model"
	"errors"
	"github.com/mitchellh/mapstructure"
)

type ProductUsecase struct {
	productRepository model.ProductRepository
}

func NewProductUsecase(pr model.ProductRepository) model.ProductUsecase {
	return ProductUsecase{
		productRepository: pr,
	}
}

func (p ProductUsecase) GetAllProduct() ([]model.Product, error) {
	return p.productRepository.GetAllProduct()
}

func (p ProductUsecase) GetProductById(id int) (model.Product, error) {
	resp, err := p.productRepository.GetProductById(id)
	if err != nil {
		return model.Product{}, errors.New("Product ID tidak ditemukan")
	}

	return resp, nil
}

func (p ProductUsecase) CreateProduct(req dto.ProductRequest) error {
	var product model.Product
	mapstructure.Decode(req, &product)

	return p.productRepository.CreateProduct(product)
}

func (p ProductUsecase) UpdateProduct(id int, req dto.ProductRequest) error {
	_, err := p.productRepository.GetProductById(id)
	if err != nil {
		return errors.New("Product ID tidak di temukan")
	}

	var product model.Product
	mapstructure.Decode(req, &product)

	return p.productRepository.UpdateProduct(id, product)
}

func (p ProductUsecase) DeleteProduct(id int) error {
	_, err := p.productRepository.GetProductById(id)
	if err != nil {
		return errors.New("Product ID tidak di temukan")
	}

	return p.productRepository.DeleteProduct(id)
}
