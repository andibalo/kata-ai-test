package model

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID             string `bun:",pk"`
	Name           string
	Email          string
	LastAccessedAt time.Time `bun:",nullzero,default:now()"`
	CreatedBy      string
	CreatedAt      time.Time `bun:",nullzero,default:now()"`
	UpdatedBy      *string
	UpdatedAt      bun.NullTime
	DeletedBy      *string
	DeletedAt      time.Time `bun:",nullzero,soft_delete"`
}
