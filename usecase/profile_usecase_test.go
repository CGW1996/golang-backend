package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/CGW1996/golang-backend/domain"
	"github.com/CGW1996/golang-backend/domain/mocks"
	"github.com/CGW1996/golang-backend/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_profileUsecase_GetProfileByID(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	userObjectID := primitive.NewObjectID()
	userID := userObjectID.Hex()

	t.Run("success", func(t *testing.T) {
		mockUser := domain.User{
			ID:    userObjectID,
			Name:  "Steven",
			Email: "s@g.c",
		}

		mockProfile := domain.Profile{
			Name:  "Steven",
			Email: "s@g.c",
		}

		mockUserRepository.On("GetByID", mock.Anything, userID).Return(mockUser, nil).Once()

		u := usecase.NewProfileUsecase(mockUserRepository, time.Second*2)
		profile, err := u.GetProfileByID(context.Background(), userID)

		assert.NoError(t, err)
		assert.NotNil(t, profile)

		if assert.NotNil(t, profile) {
			assert.Equal(t, mockProfile.Name, profile.Name)
			assert.Equal(t, mockProfile.Email, profile.Email)
		}

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepository.On("GetByID", mock.Anything, userID).Return(domain.User{}, errors.New("unexpected")).Once()
		u := usecase.NewProfileUsecase(mockUserRepository, time.Second*2)
		user, err := u.GetProfileByID(context.Background(), userID)

		assert.Error(t, err)
		assert.Nil(t, user)

		mockUserRepository.AssertExpectations(t)
	})
}
