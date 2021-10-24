package items

import (
	"golang-ecommerce-practice/errs"
	"golang-ecommerce-practice/package/entities"
)

type RepoItems interface {
	FindBy(*map[string]interface{}) (*[]entities.Items, *errs.AppError)
	FindOne(*map[string]interface{}) (*entities.Items, *errs.AppError)
	Create(*entities.Items) (*entities.Items, *errs.AppError)
	Update(*map[string]interface{}, *string) *errs.AppError
}
