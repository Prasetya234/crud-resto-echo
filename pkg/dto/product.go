package dto

import validation "github.com/go-ozzo/ozzo-validation"

type ProductRequest struct {
	ImageUrl    string `json:"image_url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

func (p ProductRequest) Validation() error {
	err := validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.Price, validation.Required),
	)
	if err != nil {
		return err
	}
	return nil
}
