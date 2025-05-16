package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Tag struct {
	bun.BaseModel `bun:"table:tags"`
	ID            int64     `bun:",pk,autoincrement"`
	UserID        int64     `bun:",notnull"`
	User          *User     `bun:"rel:belongs-to,join:user_id=id"`
	Name          string    `bun:",notnull"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Links         []*Link   `bun:"m2m:link_tags,join:Tag=TagID,join:Link=LinkID"`
}
