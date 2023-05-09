package route

import (
	"log"
	"time"

	"github.com/CGW1996/golang-backend/bootstrap"
	"github.com/CGW1996/golang-backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, routerV1 *gin.RouterGroup) {
	log.Println(routerV1)
}
