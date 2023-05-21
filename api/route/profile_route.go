package route

import (
	"time"

	"github.com/CGW1996/golang-backend/bootstrap"
	"github.com/CGW1996/golang-backend/controller"
	"github.com/CGW1996/golang-backend/domain"
	"github.com/CGW1996/golang-backend/mongo"
	"github.com/CGW1996/golang-backend/repository"
	"github.com/CGW1996/golang-backend/usecase"
	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
