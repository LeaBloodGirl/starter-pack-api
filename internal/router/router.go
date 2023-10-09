/*
The goal of the common package is to provide simple functions that will be used by other packages
Here we'll use mux for Golang
It's a package that serves to route easily requests
*/
package router

import (
	"starter-pack-api/api"
	"starter-pack-api/internal/config"
	usersorm "starter-pack-api/internal/database/users"
	"starter-pack-api/internal/logger"
	"starter-pack-api/internal/router/middleware"

	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router(logDebug, logProd logger.Logger, conf config.Config) *mux.Router {
	logginDataObject := middleware.LoggingData{LogDebugStored: &logDebug, LogStatStored: &logProd}
	r := mux.NewRouter()
	r.Use(logginDataObject.Middleware)
	usersorm.Connect(conf.UserDB)
	r.HandleFunc("/user", api.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", api.GetUserById(logDebug)).Methods("GET")
	r.HandleFunc("/user", api.CreateUser(logDebug)).Methods("POST")
	r.HandleFunc("/user/{id}", api.UpdateStatus(logDebug)).Methods("PUT")
	r.HandleFunc("/user/{id}", api.DeleteUser(logDebug)).Methods("DELETE")
	r.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)
	fs := http.FileServer(http.Dir("./docs/"))
	r.PathPrefix("/json/").Handler(http.StripPrefix("/json/", fs))
	return r
}
