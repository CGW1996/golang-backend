package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CGW1996/golang-backend/controller"
	"github.com/CGW1996/golang-backend/domain"
	"github.com/CGW1996/golang-backend/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setUserID(userID string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("x-user-id", userID)
		ctx.Next()
	}
}

func TestProfileController_Fetch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockProfile := &domain.Profile{
			Name:  "Steven",
			Email: "s@g.c",
		}

		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockProfileUsecase := new(mocks.ProfileUsecase)
		mockProfileUsecase.On("GetProfileByID", mock.Anything, userID).Return(mockProfile, nil)

		router := gin.Default()

		rec := httptest.NewRecorder()

		pc := &controller.ProfileController{
			ProfileUsecase: mockProfileUsecase,
		}

		router.Use(setUserID(userID))
		router.GET("/profile", pc.Fetch)

		body, err := json.Marshal(mockProfile)
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, bodyString, rec.Body.String())

		mockProfileUsecase.AssertExpectations(t)
	})
	t.Run("Error", func(t *testing.T) {
		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockProfileUsecase := new(mocks.ProfileUsecase)
		customErr := errors.New("unexpected")
		mockProfileUsecase.On("GetProfileByID", mock.Anything, userID).Return(nil, customErr)
		router := gin.Default()
		rec := httptest.NewRecorder()
		pc := &controller.ProfileController{
			ProfileUsecase: mockProfileUsecase,
		}
		router.Use(setUserID(userID))
		router.GET("/profile", pc.Fetch)

		body, err := json.Marshal(domain.ErrorResponse{Message: customErr.Error()})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, bodyString, rec.Body.String())

		mockProfileUsecase.AssertExpectations(t)
	})
}
