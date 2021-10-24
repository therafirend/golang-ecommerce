package entities

import (
	"golang-ecommerce-practice/errs"
	"net/http"
)

type BodyLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ChangePassword struct {
	ID           string
	OldPass      string `json:"oldPass"`
	NewPass      string `json:"password"`
	PasswordConf string `json:"passwordConf"`
}

// validation login
func (lgn *BodyLogin) ValidateLogin() *errs.AppError {
	if lgn.Username == "" {
		return errs.NewAppError("Username cannot be empty", http.StatusBadRequest)
	}

	if lgn.Password == "" {
		return errs.NewAppError("Password cannot be empty", http.StatusBadRequest)
	}

	return nil
}

// validate change password
func (cp *ChangePassword) ValidateChangePassword() *errs.AppError {
	if cp.OldPass == "" {
		return errs.NewAppError("Old Password cannot be empty", http.StatusBadRequest)
	}

	if cp.NewPass == "" {
		return errs.NewAppError("Password cannot be empty", http.StatusBadRequest)
	}

	if cp.PasswordConf == "" {
		return errs.NewAppError("Password Confirmation cannot be empty", http.StatusBadRequest)
	}

	if cp.NewPass != cp.PasswordConf {
		return errs.NewAppError("New Password and Confirm Password did not match", http.StatusBadRequest)
	}

	return nil
}
