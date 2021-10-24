package users

import (
	"golang-ecommerce-practice/errs"
	"golang-ecommerce-practice/package/entities"
)

type RepoUsers interface {
	FindBy(*map[string]interface{}) (*[]entities.Users, *errs.AppError)
	FindOne(*map[string]interface{}) (*entities.Users, *errs.AppError)
	Create(*entities.Users) (*entities.Users, *errs.AppError)
	Update(*map[string]interface{}, *string) *errs.AppError
}
