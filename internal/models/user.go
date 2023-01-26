package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expiretime"`
}
