package usecase

import (
	"context"
	"time"

	"github.com/CGW1996/golang-backend/domain"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}
