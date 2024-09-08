package repository

import (
	"github.com/uptrace/bun"
	"pokemon-be/internal/model"
)

type UserRepository interface {
	Save(user *model.User) error
	SaveTx(user *model.User, tx bun.Tx) error
	GetByEmail(email string) (*model.User, error)
	UpdateUserLastAccessedAt(email string) error
}
