package auth

import (
	"golang-ecommerce-practice/errs"
	"golang-ecommerce-practice/package/entities"
	"golang-ecommerce-practice/package/users"
	"net/http"
	"strings"
)

type RepoServiceAuth interface {
	Login(login *entities.BodyLogin) (*entities.ResponseLogin, *errs.AppError)
	ChangePassword(pw *entities.ChangePassword) *errs.AppError
	Auth(authorization *string) (ID *string, error *errs.AppError)
}

type service struct {
	repoUser users.RepoUsers
}

//NewService is used to create a single instance of the service
func NewService(repo users.RepoUsers) RepoServiceAuth {
	return &service{
		repoUser: repo,
	}
}

func (s *service) Login(login *entities.BodyLogin) (*entities.ResponseLogin, *errs.AppError) {

	if err := login.ValidateLogin(); err != nil {
		return nil, err
	}

	user, err := s.repoUser.FindOne(&map[string]interface{}{
		"username": login.Username,
		"password": entities.PasswordEncrypt(&login.Password),
		"status":   "on",
	})
	if err != nil {
		return nil, errs.NewAppError("Username or Password did not match", http.StatusBadRequest)
	}

	token, err := entities.CreateJwt(user.ToUserToken())
	if err != nil {
		return nil, err
	}

	return &entities.ResponseLogin{
		Message: "Token Berhasil Dibuat",
		Token:   *token,
	}, nil
}

func (s *service) Auth(authorization *string) (ID *string, error *errs.AppError) {
	token := strings.Split(*authorization, " ")[1]
	if token == "" {
		return nil, errs.NewAppError("Please login", http.StatusUnauthorized)
	}

	userJwt, err := entities.ValidateJwt(&token)
	if err != nil {
		return nil, err
	}

	if _, err := s.repoUser.FindOne(&map[string]interface{}{
		"id":     userJwt.ID,
		"status": "on",
	}); err != nil {
		return nil, err
	}

	return &userJwt.ID, nil
}

func (s *service) ChangePassword(pw *entities.ChangePassword) *errs.AppError {

	if err := pw.ValidateChangePassword(); err != nil {
		return err
	}

	//check user
	_, err := s.repoUser.FindOne(&map[string]interface{}{
		"id":       pw.ID,
		"password": entities.PasswordEncrypt(&pw.OldPass),
		"status":   "on",
	})
	if err != nil {
		return err
	}

	if err = s.repoUser.Update(&map[string]interface{}{
		"password": entities.PasswordEncrypt(&pw.NewPass),
	}, &pw.ID); err != nil {
		return err
	}

	return nil
}
