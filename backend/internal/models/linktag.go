package models

import "github.com/uptrace/bun"

type LinkTag struct {
	bun.BaseModel `bun:"table:link_tags"`
	LinkID        int64 `bun:"link_id,pk,notnull"`
	TagID         int64 `bun:"tag_id,pk,notnull"`
	Link          *Link `bun:"rel:belongs-to,join:link_id=id"`
	Tag           *Tag  `bun:"rel:belongs-to,join:tag_id=id"`
}
