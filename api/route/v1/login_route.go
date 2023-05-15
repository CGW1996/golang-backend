package route

import (
	"time"

	"github.com/CGW1996/golang-backend/controller"
	"github.com/CGW1996/golang-backend/domain"
	"github.com/CGW1996/golang-backend/mongo"
	"github.com/CGW1996/golang-backend/repository"
	"github.com/CGW1996/golang-backend/usecase"
	"github.com/gin-gonic/gin"
)

func NewLoginRouter(timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
	}
	group.POST("/login", lc.Login)
}
