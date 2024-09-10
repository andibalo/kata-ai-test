package repository

import (
	"context"
	"github.com/uptrace/bun"
	"pokemon-be/internal/model"
	"time"
)

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) SaveTx(user *model.User, tx bun.Tx) error {

	_, err := tx.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Save(user *model.User) error {

	_, err := r.db.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}

	err := r.db.NewSelect().Model(user).Where("email = ?", email).Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByChannelUsernameAndType(channelUsername, channelType string) (*model.User, error) {
	user := &model.User{}

	err := r.db.NewSelect().Model(user).
		Where("channel_username = ?", channelUsername).
		Where("channel_type = ?", channelType).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdateUserLastAccessedAtByChannelUsernameAndType(channelUsername, channelType string) error {
	user := &model.User{}
	user.LastAccessedAt = time.Now()

	_, err := r.db.NewUpdate().
		Model(user).
		Column("last_accessed_at").
		Where("channel_username = ?", channelUsername).
		Where("channel_type = ?", channelType).
		Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}
