package entity

import (
	"time"
)

type Users struct {
	UserID    string    `gorm:"primaryKey" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Outlet    []Outlets  `gorm:"ForeignKey:UserID"`
}

type UserInputs struct {
	UserInputID string `json:"id"`
	FullName    string `json:"full_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UpdateUserInputs struct {
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginUserInputs struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
