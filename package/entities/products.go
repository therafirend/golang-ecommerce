package entities

import (
	"golang-ecommerce-practice/errs"
	"net/http"
)

type Products struct {
	ID          string  `db:"id"`
	Name        *string `db:"name"`
	Price       *string `db:"price"`
	Stock       *string `db:"stock"`
	Description *string `db:"description"`
	Seller      *string `db:"id_seller"`
}
type CreateProducts struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
	Description string `json:"description"`
	Seller      string `json:"id_seller"`
}

func (prd *Products) ToItems() *Products {
	return &Products{
		ID:          prd.ID,
		Name:        prd.Name,
		Price:       prd.Price,
		Stock:       prd.Stock,
		Description: prd.Description,
		Seller:      prd.Seller,
	}
}
func (prd *CreateProducts) ValidateInsert() *errs.AppError {
	if prd.Name == "" {
		return errs.NewAppError("Name cannot be empty", http.StatusBadRequest)
	}

	if prd.Price == "" {
		return errs.NewAppError("Price cannot be empty", http.StatusBadRequest)
	}

	if prd.Stock == "" {
		return errs.NewAppError("Stock cannot be empty", http.StatusBadRequest)
	}
	if prd.Description == "" {
		return errs.NewAppError("Description cannot be empty", http.StatusBadRequest)
	}

	if prd.Seller == "" {
		return errs.NewAppError("ID Seller cannot be empty", http.StatusBadRequest)
	}

	return nil
}
