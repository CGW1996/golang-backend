package route

import (
	"time"

	"github.com/CGW1996/golang-backend/bootstrap"
	"github.com/CGW1996/golang-backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	NewSignupRouter(timeout, db, publicRouterV1)
	NewLoginRouter(env, timeout, db, publicRouterV1)
}
