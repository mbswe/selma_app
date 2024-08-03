package controllers

import (
	"SelmaApp/models"
	"github.com/mbswe/selma/controller"
	"net/http"
)

// UserController handles user-related requests
type UserController struct {
	controller.ControllerBase
}

// NewUserController creates a new UserController instance
func NewUserController() *UserController {
	return &UserController{
		ControllerBase: *controller.NewControllerBase(),
	}
}

// GetUser handles the retrieval of a user
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// Example action
	user := models.NewUser(1, "John Doe", "john.doe@example.com")
	w.Write([]byte("User: " + user.Name))
}
