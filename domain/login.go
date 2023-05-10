package domain

import (
	"context"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
}