package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Tweet struct {
	ID        string    `json:"id" bun:"type:uuid"`
	AuthorID  string    `json:"authorId" bun:"type:uuid"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt" bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`

	bun.BaseModel `bun:"tweets,alias:tweets"`
}
