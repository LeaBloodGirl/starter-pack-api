package main

import (
	"flag"
	"net/http"
	"starter-pack-api/internal/config"
	"starter-pack-api/internal/logger"
	"starter-pack-api/internal/router"
	"strconv"
	"time"
)

// @title API TISSEO GO
// @version 1.0
// @description This is the API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	//Defining launch flags
	filepath := flag.String("configpath", "", "Give the path to the config file")
	//Getting flags in their respective variables
	flag.Parse()

	//Instantiating the config
	config.Get(*filepath)

	//Instantiating the loggers
	logDebug := logger.OpenLogger(config.ConfigLoaded.Logs1)
	logProd := logger.OpenLogger(config.ConfigLoaded.Logs2)

	//Instantiating the router
	r := router.Router(logDebug, logProd, config.ConfigLoaded)

	//Instantiating the server
	srv := &http.Server{
		Handler: r,
		Addr:    config.ConfigLoaded.Server.Host + ":" + strconv.Itoa(config.ConfigLoaded.Server.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	//Starting the server in https if there's a certificate (it should always be this way)
	if config.ConfigLoaded.Server.Certif == "" || config.ConfigLoaded.Server.CertifKey == "" {
		srv.ListenAndServe()
	} else { //or in http if there isn't a certificate
		srv.ListenAndServeTLS(config.ConfigLoaded.Server.Certif, config.ConfigLoaded.Server.CertifKey)
	}
}
