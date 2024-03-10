package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type JwtCustomClaims struct {
	ID uuid.UUID `json:"name"`
	jwt.RegisteredClaims
}

type Account struct {
	ID       uuid.UUID
	Email    string
	Password string
}

type Group struct {
	ID   uuid.UUID
	Name string
}

type Todo struct {
	ID          uuid.UUID
	Title       string
	description string
	isCompleted bool
	dueDate     time.Time
	GroupID     uuid.UUID
}

type UserGroup struct {
	UserID  uuid.UUID
	GroupID uuid.UUID
}
