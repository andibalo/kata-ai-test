package repository

import (
	"github.com/uptrace/bun"
	"pokemon-be/internal/model"
)

type UserRepository interface {
	Save(user *model.User) error
	SaveTx(user *model.User, tx bun.Tx) error
	GetByEmail(email string) (*model.User, error)
	GetByChannelUsernameAndType(channelUsername, channelType string) (*model.User, error)
	UpdateUserLastAccessedAtByChannelUsernameAndType(channelUsername, channelType string) error
}
