package service

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"pokemon-be/internal/apperr"
	"pokemon-be/internal/config"
	"pokemon-be/internal/model"
	"pokemon-be/internal/repository"
	"pokemon-be/internal/request"
	"time"
)

type userService struct {
	cfg      config.Config
	userRepo repository.UserRepository
}

func NewUserService(cfg config.Config, userRepo repository.UserRepository) UserService {

	return &userService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(data *request.RegisterUserRequest) error {

	existingUser, err := s.userRepo.GetByChannelUsernameAndType(data.ChannelUsername, data.ChannelType)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.cfg.Logger().Error().Err(err).Msg("[CreateUser] Failed to get user by channel username and type")
		return err
	}

	if existingUser != nil {
		s.cfg.Logger().Error().Err(err).Msg("[CreateUser] User already exists")
		return errors.New("User already exists")
	}

	user := &model.User{
		ID:              uuid.NewString(),
		Name:            data.Name,
		Email:           data.Email,
		ChannelUsername: data.ChannelUsername,
		ChannelType:     data.ChannelType,
		LastAccessedAt:  time.Now(),
		CreatedAt:       time.Now(),
	}

	err = s.userRepo.Save(user)
	if err != nil {
		s.cfg.Logger().Error().Err(err).Msg("[CreateUser] Failed to insert user to database")
		return err
	}

	return nil
}

func (s *userService) Login(data *request.LoginRequest) (*model.User, error) {

	existingUser, err := s.userRepo.GetByChannelUsernameAndType(data.ChannelUsername, data.ChannelType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.cfg.Logger().Error().Err(err).Msg("[Login] Invalid email/password")
			return existingUser, apperr.ErrNotFound
		}

		s.cfg.Logger().Error().Err(err).Msg("[Login] Failed to get user by channel username and type")
		return existingUser, err
	}

	err = s.userRepo.UpdateUserLastAccessedAtByChannelUsernameAndType(data.ChannelUsername, data.ChannelType)
	if err != nil {
		s.cfg.Logger().Error().Err(err).Msg("[Login] Failed to update user last accessed at")
		return existingUser, err
	}

	return existingUser, nil
}
