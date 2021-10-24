package users

import (
	"golang-ecommerce-practice/errs"
	"golang-ecommerce-practice/package/entities"
)

type RepoServiceUsers interface {
	GetUser(where *map[string]interface{}) (*entities.UserToken, *errs.AppError)
	GetUsers(where *map[string]interface{}) (*[]entities.UserToken, *errs.AppError)
	InsertUser(insert *entities.RegisUser) (*entities.UserToken, *errs.AppError)
	UpdateUser(setData *map[string]interface{}, id *string) *errs.AppError
	DeleteUser(id *string) *errs.AppError
}

type service struct {
	repoUser RepoUsers
}

func NewService(repo RepoUsers) RepoServiceUsers {
	return &service{
		repoUser: repo,
	}
}

func (srv *service) GetUser(where *map[string]interface{}) (*entities.UserToken, *errs.AppError) {
	usr, err := srv.repoUser.FindOne(where)

	if err != nil {
		return nil, err
	}

	return usr.ToUserToken(), nil
}

func (srv *service) GetUsers(where *map[string]interface{}) (*[]entities.UserToken, *errs.AppError) {
	usr, err := srv.repoUser.FindBy(where)

	if err != nil {
		return nil, err
	}

	usersToken := make([]entities.UserToken, 0)

	for _, val := range *usr {
		usersToken = append(usersToken, *val.ToUserToken())
	}

	return &usersToken, nil
}

func (srv *service) InsertUser(insert *entities.RegisUser) (*entities.UserToken, *errs.AppError) {
	if err := insert.ValidateRegis(); err != nil {
		return nil, err
	}

	usr := &entities.Users{
		ID:       insert.CreateUniqId(),
		Username: insert.Username,
		Name:     insert.Name,
		Password: entities.PasswordEncrypt(&insert.Password),
		Status:   "on",
	}

	if _, err := srv.repoUser.Create(usr); err != nil {
		return nil, err
	}

	return usr.ToUserToken(), nil
}

func (srv *service) UpdateUser(setData *map[string]interface{}, id *string) *errs.AppError {
	if _, err := srv.repoUser.FindOne(&map[string]interface{}{
		"id": *id,
	}); err != nil {
		return err
	}

	for i := range *setData {
		if i != "username" && i != "Name" {
			delete(*setData, i)
		}
	}

	if err := srv.repoUser.Update(setData, id); err != nil {
		return err
	}

	return nil
}

func (srv *service) DeleteUser(id *string) *errs.AppError {
	if _, err := srv.repoUser.FindOne(&map[string]interface{}{
		"id": *id,
	}); err != nil {
		return err
	}

	if err := srv.repoUser.Update(&map[string]interface{}{"status": "off"}, id); err != nil {
		return err
	}

	return nil
}
