package products

import (
	"golang-ecommerce-practice/errs"
	"golang-ecommerce-practice/package/entities"
)

type RepoProducts interface {
	FindBy(*map[string]interface{}) (*[]entities.Products, *errs.AppError)
	FindOne(*map[string]interface{}) (*entities.Products, *errs.AppError)
	Create(*entities.CreateProducts) (*entities.CreateProducts, *errs.AppError)
	Update(*map[string]interface{}, *string) *errs.AppError
	Delete(*map[string]interface{}) *errs.AppError
}
