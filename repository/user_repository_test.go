package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/CGW1996/golang-backend/domain"
	"github.com/CGW1996/golang-backend/mongo/mocks"
	"github.com/CGW1996/golang-backend/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_userRepository_Create(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionUser

	mockUser := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     "Steven",
		Email:    "Steven@gmail.com",
		Password: "password",
	}
	mockEmptyUser := &domain.User{}
	mockUserID := primitive.NewObjectID()

	t.Run("success", func(t *testing.T) {
		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockUserID, nil).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewUserRepository(databaseHelper, collectionName)

		err := ur.Create(context.Background(), mockUser)

		assert.NoError(t, err)
		collectionHelper.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		collectionHelper.On("InsertOne", mock.Anything, mock.AnythingOfType("*domain.User")).Return(mockEmptyUser, errors.New("unexpected")).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewUserRepository(databaseHelper, collectionName)

		err := ur.Create(context.Background(), mockUser)

		assert.Error(t, err)
		collectionHelper.AssertExpectations(t)
	})
}

func Test_userRepository_Fetch(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}
	cursorHelper := &mocks.Cursor{}

	collectionName := domain.CollectionUser

	mockUser := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     "Steven",
		Email:    "Steven@gmail.com",
		Password: "password",
	}

	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, *mockUser)

	t.Run("success", func(t *testing.T) {
		collectionHelper.On("Find", mock.Anything, mock.AnythingOfType("primitive.D"), mock.Anything).Return(cursorHelper, nil).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper).Once()
		cursorHelper.On("All", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			arg := args.Get(1).(*[]domain.User)
			*arg = append(*arg, *mockUser)
		})
		ur := repository.NewUserRepository(databaseHelper, collectionName)
		users, err := ur.Fetch(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Len(t, users, len(mockListUser))
	})
}
