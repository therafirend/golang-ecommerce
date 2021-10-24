package items

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

func (rp repoDb) FindBy(where *map[string]interface{}) (*[]entities.Items, *errs.AppError) {
	items := make([]entities.Items, 0)
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
	findSql := "SELECT * FROM items " + whereSql

	// logs.info(findSql)
	err := rp.client.Select(&items, findSql, value...)
	if err != nil {
		zapLog.Error("Error sql " + err.Error())
		return nil, errs.NewAppError("Error", http.StatusInternalServerError)
	}

	return &items, nil
}

func (rp repoDb) FindOne(where *map[string]interface{}) (*entities.Items, *errs.AppError) {
	items, err := rp.FindBy(where)

	if err != nil {
		return nil, errs.NewAppError("Error", http.StatusInternalServerError)
	}

	res := *items

	if len(res) <= 0 {
		return nil, errs.NewAppError("User not found", http.StatusNotFound)
	}

	return &res[0], nil
}

func (p repoDb) Create(where *entities.Items) (*entities.Items, *errs.AppError) {
	panic("implement me")
}

func (p repoDb) Update(where *map[string]interface{}, id *string) *errs.AppError {
	panic("implement me")
}

func NewRepoDB(client *sqlx.DB) RepoItems {
	return &repoDb{
		client: client,
	}
}
