package entities


import (
	"github.com/segmentio/ksuid"
	"golang-ecommerce-practice/errs"
	"net/http"
)

type Users struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Name      string `db:"name"`
	Password string `db:"password"`
	Status   string `db:"status"`
}

type RegisUser struct {
	Username    string `json:"username"`
	Name         string `json:"name"`
	Password    string `json:"password"`
	PassConfirm string `json:"passConfirm"`
}

type UserToken struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name      string `json:"name"`
}

func (usr *RegisUser) CreateUniqId() string {
	return ksuid.New().String()
}

func (usr *Users) ToUserToken() *UserToken  {
	return &UserToken{
		ID: usr.ID,
		Username: usr.Username,
		Name: usr.Name,
	}
}

func (usr *RegisUser) ValidateRegis() *errs.AppError  {
	if usr.Username == "" {
		return errs.NewAppError("Username cannot be empty", http.StatusBadRequest)
	}

	if usr.Name == "" {
		return errs.NewAppError("Name cannot be empty", http.StatusBadRequest)
	}

	if usr.Password == "" {
		return errs.NewAppError("Password cannot be empty", http.StatusBadRequest)
	}

	if usr.PassConfirm == "" {
		return errs.NewAppError("Password Confirmation cannot be empty", http.StatusBadRequest)
	}

	if usr.Password !=  usr.PassConfirm {
		return errs.NewAppError("Password and Password Confirmation did not match", http.StatusBadRequest)
	}

	return nil
}
