package route

import (
	"fmt"
	"time"

	"github.com/CGW1996/golang-backend/bootstrap"
	"github.com/CGW1996/golang-backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, routerV1 *gin.RouterGroup) {
	fmt.Println(env)
	publicRouterV1 := routerV1.Group("")
	NewSignupRouter(timeout, db, publicRouterV1)
}
