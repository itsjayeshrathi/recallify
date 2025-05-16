package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Link struct {
	bun.BaseModel `bun:"table:links"`
	ID            int64     `bun:",pk,autoincrement"`
	UserID        int64     `bun:",notnull"`
	User          *User     `bun:"rel:belongs-to,join:user_id=id"`
	URL           string    `bun:",notnull"`
	Slug          string    `bun:",unique,notnull"`
	Summary       string    `bun:",type:text"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Tags          []*Tag    `bun:"m2m:link_tags,join:Link=LinkID,join:Tag=TagID"`
}
