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
	userController := controllers.NewUserController(app)

	app.Router.HandleFunc("/", homeController.ShowIndex).Methods("GET")
	app.Router.HandleFunc("/user/{id}", userController.GetUser).Methods("GET")
	app.Router.HandleFunc("/user", userController.SaveUser).Methods("POST")

	app.Router.Use(selmaMiddleware.LoggingMiddleware(app))
	app.Router.Use(appMiddleware.CustomMiddleware(app))

	app.StartServer()
}
