package controllers

import (
	"SelmaApp/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mbswe/selma"
	"log"
	"net/http"
	"time"
)

// UserController handles user-related requests
type UserController struct {
	App *selma.App
}

// NewUserController creates a new UserController instance
func NewUserController(app *selma.App) *UserController {
	return &UserController{
		App: app,
	}
}

// SaveUser handles saving a user to the database
func (uc *UserController) SaveUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Ensure all mandatory fields are present
	if user.Email == "" || user.Username == "" || user.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Set timestamps for create and update
	currentTime := time.Now()
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	// Save user to the database using GORM
	if err := uc.App.DB.Save(&user).Error; err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		log.Printf("Failed to save user: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// GetUser retrieves a user by ID
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	// Find the user in the database
	var user models.User
	if err := uc.App.DB.First(&user, id).Error; err != nil {
		http.Error(w, "Failed to find user", http.StatusNotFound)
		log.Printf("Failed to find user: %v", err)
		return
	}

	// Return the user as JSON
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
