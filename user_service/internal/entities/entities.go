package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	ID        string    `json:"id" bun:"type:uuid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     *string   `json:"email"` // TODO: why is optional?
	CreatedAt time.Time `json:"registeredAt" bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`
}

type FullUserInfo struct {
	User
	Password string `json:"-"`

	bun.BaseModel `bun:"users,alias:users"`
}
