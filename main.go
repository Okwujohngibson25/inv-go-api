package main

import (
	"github.com/gin-gonic/gin"

	"github.com/coinserveringo/config"
	_ "github.com/coinserveringo/docs"
	"github.com/coinserveringo/internal/app"
	"github.com/coinserveringo/internal/routes"
	"github.com/coinserveringo/mail"
)

// @title       My Inv API
// @version     1.0
// @description This is my Go Inv API.
// @host        localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	router := gin.Default()
	configdata, _ := config.LoadConfig()
	dbSqlc := config.ConnectDB(configdata)
	mailer := mail.NewMailService(configdata)
	myApp := app.NewApp(dbSqlc, mailer, configdata)
	routes.RegisterRoutes(router, myApp)

	// Start server
	router.Run("localhost:8080")
}
