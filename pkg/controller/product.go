package controller

import (
	"crud_product/pkg/dto"
	"crud_product/pkg/model"
	"crud_product/shared/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductController struct {
	ProductUsecase model.ProductUsecase
}

func NewProductController(pu model.ProductUsecase) ProductController {
	return ProductController{
		ProductUsecase: pu,
	}
}

func (p *ProductController) GetAllProduct(c echo.Context) error {

	resp, err := p.ProductUsecase.GetAllProduct()
	if err != nil {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.SetResponse(c, http.StatusOK, "success", resp)
}

func (p *ProductController) GetProductById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := p.ProductUsecase.GetProductById(id)
	if err != nil {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.SetResponse(c, http.StatusOK, "success", resp)
}

func (p *ProductController) CreateProduct(c echo.Context) error {
	var request dto.ProductRequest
	c.Bind(&request)

	err := request.Validation()
	if err != nil {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err = p.ProductUsecase.CreateProduct(request)
	if err != nil {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.SetResponse(c, http.StatusOK, "success", nil)
}

func (p *ProductController) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var request dto.ProductRequest
	c.Bind(&request)

	err := request.Validation()
	if err != nil {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err = p.ProductUsecase.UpdateProduct(id, request)
	if err != nil {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.SetResponse(c, http.StatusOK, "success", nil)
}

func (p *ProductController) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := p.ProductUsecase.DeleteProduct(id)
	if err != nil {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return utils.SetResponse(c, http.StatusOK, "success", nil)
}
