package products

import (
	"github.com/jmoiron/sqlx"
	"golang-ecommerce-practice/errs"
	"golang-ecommerce-practice/package/entities"
	"golang-ecommerce-practice/zapLog"
	"net/http"
	"reflect"
	"strings"
)

type repoDb struct {
	client *sqlx.DB
}

func (rp repoDb) FindBy(where *map[string]interface{}) (*[]entities.Products, *errs.AppError) {
	items := make([]entities.Products, 0)
	whereSql := ""
	value := make([]interface{}, 0)
	for key, val := range *where {
		if whereSql == "" {
			whereSql += " WHERE "
		} else {
			whereSql += " AND "
		}
		tmp := ""

		if reflect.TypeOf(val).String() == "string" {
			s := strings.Split(val.(string), " ")
			if len(s) > 1 {
				f := strings.ToLower(s[0])
				if f == "between" || f == "like" {
					tmp = key + " " + val.(string)
				}
			}
		}
		if tmp == "" {
			tmp = key + "=?"
			value = append(value, val)
		}

		whereSql += tmp
	}
	findSql := "SELECT * FROM products " + whereSql

	// logs.info(findSql)
	err := rp.client.Select(&items, findSql, value...)
	if err != nil {
		zapLog.Error("Error sql " + err.Error())
		return nil, errs.NewAppError("Error", http.StatusInternalServerError)
	}

	return &items, nil
}

func (rp repoDb) FindOne(where *map[string]interface{}) (*entities.Products, *errs.AppError) {
	items, err := rp.FindBy(where)

	if err != nil {
		return nil, errs.NewAppError("Error", http.StatusInternalServerError)
	}

	res := *items

	if len(res) <= 0 {
		return nil, errs.NewAppError("Product not found", http.StatusNotFound)
	}

	return &res[0], nil
}

func (rp repoDb) Create(products *entities.CreateProducts) (*entities.CreateProducts, *errs.AppError) {
	insertSql := "insert into products(`name`, `price`, `stock`, `description`, `id_seller`) values(?,?,?,?,?)"
	_, err := rp.client.Exec(insertSql, products.Name, products.Price, products.Stock, products.Description, products.Seller)

	if err != nil {
		zapLog.Error("Error Sql" + err.Error())

		if strings.Contains(err.Error(), "Duplicate Entry") {
			return nil, errs.NewAppError("Product Already Exist", http.StatusBadRequest)
		}

		return nil, errs.NewAppError("Internal server error", http.StatusInternalServerError)
	}

	return products, nil
}

func (rp repoDb) Update(setData *map[string]interface{}, id *string) *errs.AppError {
	setSql := ""
	value := make([]interface{}, 0)

	for key, val := range *setData {
		if setSql != "" {
			setSql += ", "
		}

		value = append(value, val)
		setSql += key + "=?"
	}

	value = append(value, *id)

	updateSql := "update products set " + setSql + " where id = ?"
	_, err := rp.client.Exec(updateSql, value...)

	if err != nil {
		zapLog.Info("Error sql" + err.Error())
		return errs.NewAppError("Error", http.StatusInternalServerError)
	}

	return nil
}

func (rp repoDb) Delete(where *map[string]interface{}) *errs.AppError {
	whereSQL := ""
	value := make([]interface{}, 0)
	for k, v := range *where {
		if whereSQL == "" {
			whereSQL += " WHERE "
		} else {
			whereSQL += " AND "
		}
		whereSQL += k + "=?"
		value = append(value, v)
	}
	deleteSQL := "delete from products " + whereSQL
	_, err := rp.client.Exec(deleteSQL, value...)
	if err != nil {
		zapLog.Error("Error Delete " + err.Error())
		return errs.NewAppError("Error", http.StatusInternalServerError)
	}
	return nil
}

func NewRepoDB(client *sqlx.DB) RepoProducts {
	return &repoDb{
		client: client,
	}
}
