package users

import (
	"encoding/json"
	"errors"
	"net/http"
	usersorm "starter-pack-api/internal/database/users"
	"starter-pack-api/internal/logger"
	"starter-pack-api/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const contentType = "Content-Type"
const applicationJson = "application/json"
const userNotFound = "User Not Found!"

type PostNewUser struct {
	Email       string `json:"email"`
	Login       string `json:"login"`
	Application string `json:"app_name"`
}

type PostStatus struct {
	ValidUntil string `json:"valid_until"`
	AppName    string `json:"app_name"`
}

func checkIfUserExists(userId string) bool {
	var user models.User
	usersorm.Instance.First(&user, userId)
	if user.ID == 0 {
		return false
	}
	return true
}

// Response
type UserResponse struct {
	models.User
}

// @Summary Create User
// @Id 1
// @version 1.0
// @produce application/json
// @Param email body string true "email of the new user"
// @Param login body string true "login of the new user"
// @Param app_name body string true "the application authorized for the user"
// @Success 200 {object} UserResponse
// @Router /user [post]
func CreateUser(l logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJson)

		var postData PostNewUser
		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := models.User{
			Email:     postData.Email,
			Login:     postData.Login,
			AppName:   postData.Application,
			Token:     uuid.New().String(),
			CreatedAt: time.Now(),
		}
		response := UserResponse{
			User: user,
		}
		usersorm.Instance.Create(&user)
		json.NewEncoder(w).Encode(response)
	}
}

// @Summary Get Users
// @Id 1
// @version 1.0
// @produce application/json
// @Success 200 {object} []UserResponse
// @Router /user [get]
func GetUsers() []models.User {
	var users []models.User
	usersorm.Instance.Find(&users)
	return users
}

// @Summary This is the function to get a user by it's id
// @Description This is the description for getting a thing by its UUID. Which can be longer,
// @Param id path string true "The ID of a user"
// @Success 200 {object} UserResponse
// @Router /user/{id} [get]
func GetUserById(l logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["id"]
		if checkIfUserExists(userId) == false {
			json.NewEncoder(w).Encode(userNotFound)
			l.InfoLevel("User not found for get")
			return
		}
		var user models.User
		usersorm.Instance.First(&user, userId)
		w.Header().Set(contentType, applicationJson)
		json.NewEncoder(w).Encode(UserResponse{User: user})
	}
}

// @Summary This is the function to update a user by it's id
// @Description This is the description for getting a thing by its UUID. Which can be longer,
// @Param id path string true "The ID of a user"
// @Param valid_until body string true "The date until token expiration"
// @Param app_name body string true "The application for the token"
// @Success 200 {object} UserResponse
// @Router /user/{id} [put]
func UpdateStatus(l logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["id"]
		if checkIfUserExists(userId) == false {
			http.Error(w, userNotFound, http.StatusInternalServerError)
			l.ErrorLevel("User Not Found : ", errors.New("user was not found in the database"))
			return
		}
		var user models.User
		usersorm.Instance.First(&user, userId)
		var postedData PostStatus
		err := json.NewDecoder(r.Body).Decode(&postedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			l.ErrorLevel("Error while decoding posted body : ", err)
			return
		}
		parsedTime, err := time.Parse("2006-01-02", postedData.ValidUntil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			l.ErrorLevel("Error while parsing datetime : ", err)
			return
		}
		user.ValidUntil = parsedTime
		usersorm.Instance.Save(&user)
		w.Header().Set(contentType, applicationJson)
		json.NewEncoder(w).Encode(UserResponse{User: user})
	}
}

// @Summary This is the function to delete a user by it's id
// @Description This is the description for getting a thing by its UUID. Which can be longer,
// @Param id path string true "The ID of a user"
// @Success 200 {string} string "User Deleted Successfully!"
// @Router /user/{id} [delete]
func DeleteUser(l logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationJson)
		userId := mux.Vars(r)["id"]
		if checkIfUserExists(userId) == false {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(userNotFound)
			l.InfoLevel("No user found for deletion")
			return
		}
		var user models.User
		usersorm.Instance.Delete(&user, userId)
		json.NewEncoder(w).Encode("User Deleted Successfully!")
	}
}
