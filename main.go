package main

import (
	"SelmaApp/controllers"
	appMiddleware "SelmaApp/middlewares"
	"github.com/mbswe/selma"
	selmaMiddleware "github.com/mbswe/selma/middleware"
)

func main() {
	app := selma.NewApp("config.json")

	app.RunMigrations()
	
	homeController := controllers.NewHomeController(app)
	userController := controllers.NewUserController()

	app.Router.Get("/", homeController.ShowIndex, selmaMiddleware.LoggingMiddleware(app))
	app.Router.Get("/user", userController.GetUser, selmaMiddleware.LoggingMiddleware(app), appMiddleware.CustomMiddleware(app), selmaMiddleware.AuthMiddleware(app))

	app.StartServer()
}
