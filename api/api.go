package api

import (
	"encoding/json"
	"net/http"
	"starter-pack-api/internal/app/users"
	"starter-pack-api/internal/logger"
	"starter-pack-api/internal/models"
)

// Response
type UserResponse struct {
	models.User
}

// @Summary Get Users
// @version 1.0
// @produce application/json
// @Success 200 {array} models.User
// @Router /user [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := users.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// @Summary Create User
// @version 1.0
// @produce application/json
// @Param email body string true "email of the new user"
// @Param login body string true "login of the new user"
// @Param app_name body string true "the application authorized for the user"
// @Success 200 {object} models.User
// @Router /user [post]
func CreateUser(logDebug logger.Logger) http.HandlerFunc {
	return users.CreateUser(logDebug)
}

// @Summary This is the function to get a user by it's id
// @Description This is the description for getting a thing by its UUID. Which can be longer,
// @Param id path string true "The ID of a user"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func GetUserById(logDebug logger.Logger) http.HandlerFunc {
	return users.GetUserById(logDebug)
}

// @Summary This is the function to update a user by it's id
// @Description This is the description for getting a thing by its UUID. Which can be longer,
// @Param id path string true "The ID of a user"
// @Param valid_until body string true "The date until token expiration"
// @Param app_name body string true "The application for the token"
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func UpdateStatus(logDebug logger.Logger) http.HandlerFunc {
	return users.UpdateStatus(logDebug)
}

// @Summary This is the function to delete a user by it's id
// @Description This is the description for getting a thing by its UUID. Which can be longer,
// @Param id path string true "The ID of a user"
// @Success 200 {string} string "User Deleted Successfully!"
// @Router /user/{id} [delete]
func DeleteUser(logDebug logger.Logger) http.HandlerFunc {
	return users.DeleteUser(logDebug)
}
