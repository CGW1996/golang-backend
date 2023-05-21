package main

import (
	"time"

	route "github.com/CGW1996/golang-backend/api/route"
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

	route.Setup(env, timeout, db, router)

	router.Run(env.ServerAddress)
}
