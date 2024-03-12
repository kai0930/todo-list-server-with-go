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
	ID       uuid.UUID `json:"id" gorm:"primary_key"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type Group struct {
	ID   uuid.UUID `json:"id" gorm:"primary_key"`
	Name string    `json:"name"`
}

type Todo struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	DueDate     time.Time `json:"dueDate"`
	GroupID     uuid.UUID `json:"groupID"`
}

type UserGroup struct {
	UserID  uuid.UUID `json:"userID" gorm:"foreignKey:UserID:references:ID"`
	GroupID uuid.UUID `json:"groupID" gorm:"foreignKey:GroupID:references:ID"`
}
type GroupRequest struct {
	Name *string `json:"name"`
}

type CreateTodoRequest struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	IsCompleted *bool      `json:"isCompleted"`
	DueDate     *time.Time `json:"dueDate"`
	GroupID     uuid.UUID  `json:"groupID"`
}

type PutTodoRequest struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	IsCompleted *bool      `json:"isCompleted"`
	DueDate     *time.Time `json:"dueDate"`
	GroupID     *uuid.UUID `json:"groupID"`
}
