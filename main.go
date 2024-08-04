package main

import (
	"SelmaApp/controllers"
	appMiddleware "SelmaApp/middlewares"
	"SelmaApp/models"
	"github.com/mbswe/selma"
	selmaMiddleware "github.com/mbswe/selma/middleware"
)

func main() {
	app := selma.NewApp("config.json")

	app.RunMigrations(&models.User{})

	homeController := controllers.NewHomeController(app)
	userController := controllers.NewUserController(app.DB)

	app.Router.Get("/", homeController.ShowIndex, selmaMiddleware.LoggingMiddleware(app))
	// app.Router.Get("/user", userController.GetUser, selmaMiddleware.LoggingMiddleware(app), appMiddleware.CustomMiddleware(app) /*, selmaMiddleware.AuthMiddleware(app)*/)
	app.Router.Post("/user", userController.SaveUser, selmaMiddleware.LoggingMiddleware(app), appMiddleware.CustomMiddleware(app) /*, selmaMiddleware.AuthMiddleware(app)*/)

	app.StartServer()
}
