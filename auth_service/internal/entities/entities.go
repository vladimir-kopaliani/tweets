package entities

import (
	"time"
)

type User struct {
	ID        string    `json:"id" bun:"type:uuid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"registeredAt" bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`
}

type FullUserInfo struct {
	User
	Password string `json:"-"`
}
