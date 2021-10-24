package users

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

func (r repoDb) FindBy(where *map[string]interface{}) (*[]entities.Users, *errs.AppError) {
	items := make([]entities.Users, 0)
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
	findSql := "select * from users " + whereSql

	// logs.info(findSql)
	err := r.client.Select(&items, findSql, value...)
	if err != nil {
		zapLog.Error("Error sql " + err.Error())
		return nil, errs.NewAppError("Error", http.StatusInternalServerError)
	}

	return &items, nil
}

func (r repoDb) FindOne(where *map[string]interface{}) (*entities.Users, *errs.AppError) {
	items, err := r.FindBy(where)

	if err != nil {
		return nil, errs.NewAppError("Error", http.StatusInternalServerError)
	}

	res := *items

	if len(res) <= 0 {
		return nil, errs.NewAppError("User not found", http.StatusNotFound)
	}

	return &res[0], nil
}

func (r repoDb) Create(users *entities.Users) (*entities.Users, *errs.AppError) {
	insertSql := "insert into users(`id`, `username`, `name`, `password`) values(?,?,?,?)"
	_, err := r.client.Exec(insertSql, users.ID, users.Username, users.Name, users.Password)

	if err != nil {
		zapLog.Error("Error Sql" + err.Error())

		if strings.Contains(err.Error(), "Duplicate Entry") {
			return nil, errs.NewAppError("Username Already Exist", http.StatusBadRequest)
		}

		return nil, errs.NewAppError("Internal server error", http.StatusInternalServerError)
	}

	return users, nil
}

func (r repoDb) Update(setData *map[string]interface{}, id *string) *errs.AppError {
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

	updateSql := "update users set " + setSql + " where id = ?"
	_, err := r.client.Exec(updateSql, value...)

	if err != nil {
		zapLog.Info("Error sql" + err.Error())
		return errs.NewAppError("Error", http.StatusInternalServerError)
	}

	return nil
}

func NewRepoDB(client *sqlx.DB) RepoUsers {
	return &repoDb{
		client: client,
	}
}
