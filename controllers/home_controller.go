package controllers

import (
	"github.com/mbswe/selma"
	"github.com/mbswe/selma/controller"
	"net/http"
)

// HomeController implementing application-specific logic
type HomeController struct {
	controller.ControllerBase
	*selma.App
}

// NewHomeController creates a new instance of HomeController
func NewHomeController(app *selma.App) *HomeController {
	return &HomeController{
		ControllerBase: *controller.NewControllerBase(),
		App:            app,
	}
}

// ShowIndex serves the HomePage
func (hc *HomeController) ShowIndex(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
		Text    string
		Names   []string
	}{
		Title:   "Home",
		Message: "Welcome to the Home Page",
		Text:    "This is the text for the Home Page",
		Names:   []string{"Alice", "Bob", "Charlie"},
	}
	err := hc.App.ViewRenderer.Render(w, r, "home.html", data)
	if err != nil {
		http.Error(w, "Failed to render view", http.StatusInternalServerError)
	}
}
