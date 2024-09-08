package service

import (
	"pokemon-be/internal/model"
	"pokemon-be/internal/request"
)

type UserService interface {
	CreateUser(data *request.RegisterUserRequest) error
	Login(data *request.LoginRequest) (*model.User, error)
}
