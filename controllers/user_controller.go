package controllers

import (
	"SelmaApp/models"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

// UserController handles user-related requests
type UserController struct {
	DB *gorm.DB
}

// NewUserController creates a new UserController instance
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
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
	if err := uc.DB.Save(&user).Error; err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		log.Printf("Failed to save user: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
