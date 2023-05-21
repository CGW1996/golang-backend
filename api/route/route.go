package route

import (
	"time"

	"github.com/CGW1996/golang-backend/api/middleware"
	"github.com/CGW1996/golang-backend/bootstrap"
	"github.com/CGW1996/golang-backend/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.Engine) {
	publicRouter := router.Group("")
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := router.Group("")
	// middleware to verify accessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// all private api
	NewProfileRouter(env, timeout, db, protectedRouter)
}
