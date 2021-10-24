package products

import (
	"golang-ecommerce-practice/errs"
	"golang-ecommerce-practice/package/entities"
)

type RepoServiceProducts interface {
	GetProduct(where *map[string]interface{}) (*entities.Products, *errs.AppError)
	GetProducts(where *map[string]interface{}) (*[]entities.Products, *errs.AppError)
	InsertProducts(insert *entities.CreateProducts) (*entities.CreateProducts, *errs.AppError)
	UpdateProducts(setData *map[string]interface{}, id *string) *errs.AppError
	DeleteProducts(where *map[string]interface{}) *errs.AppError
}
type service struct {
	repoProducts RepoProducts
}

func NewService(repo RepoProducts) RepoServiceProducts {
	return &service{
		repoProducts: repo,
	}
}

func (srv service) GetProduct(where *map[string]interface{}) (*entities.Products, *errs.AppError) {
	prd, err := srv.repoProducts.FindOne(where)
	if err != nil {
		return nil, err
	}
	return prd.ToItems(), nil
}

func (srv service) GetProducts(where *map[string]interface{}) (*[]entities.Products, *errs.AppError) {
	prd, err := srv.repoProducts.FindBy(where)

	if err != nil {
		return nil, err
	}

	products := make([]entities.Products, 0)

	for _, val := range *prd {
		products = append(products, *val.ToItems())
	}

	return &products, nil
}

func (srv service) InsertProducts(insert *entities.CreateProducts) (*entities.CreateProducts, *errs.AppError) {
	if err := insert.ValidateInsert(); err != nil {
		return nil, err
	}

	prd := &entities.CreateProducts{
		Name:        insert.Name,
		Price:       insert.Price,
		Stock:       insert.Stock,
		Description: insert.Description,
		Seller:      insert.Seller,
	}
	product, err := srv.repoProducts.Create(prd)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (srv service) UpdateProducts(setData *map[string]interface{}, id *string) *errs.AppError {
	if _, err := srv.repoProducts.FindOne(&map[string]interface{}{
		"id": *id,
	}); err != nil {
		return err
	}

	for i := range *setData {
		if i != "name" && i != "price" && i != "stock" && i != "id_seller" {
			delete(*setData, i)
		}
	}

	if err := srv.repoProducts.Update(setData, id); err != nil {
		return err
	}

	return nil
}

func (srv service) DeleteProducts(where *map[string]interface{}) *errs.AppError {
	if _, err := srv.repoProducts.FindOne(where); err != nil {
		return err
	}
	if err := srv.repoProducts.Delete(where); err != nil {
		return err
	}
	//if err := srv.repoProducts.Update(&map[string]interface{}{"id": "off"}, id); err != nil {
	//	return err
	//}

	return nil
}
