package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           int64     `bun:",pk,autoincrement"`
	Name         string    `bun:",notnull"`
	Email        string    `bun:",unique,notnull"`
	PasswordHash string    `bun:",notnull"`
	Bio          string    `bun:"type:text"`
	CreatedAt    time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
