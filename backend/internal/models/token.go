package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Token struct {
	bun.BaseModel `bun:"table:tokens"`

	Hash   []byte    `bun:",pk"`
	UserID int64     `bun:",notnull"`
	User   *User     `bun:"rel:belongs-to,join:user_id=id"`
	Expiry time.Time `bun:",notnull"`
	Scope  string    `bun:",notnull"`
}
