package main

import (
	"time"

	routeV1 "github.com/CGW1996/golang-backend/api/route/v1"
	"github.com/CGW1996/golang-backend/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		c.String(200, "/user")
	})

	routerV1 := router.Group("/v1")

	routeV1.Setup(env, timeout, db, routerV1)

	router.Run(env.ServerAddress)
}
