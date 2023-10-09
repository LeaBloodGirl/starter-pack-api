/*
	The goal of the config package is to load a config file written in yaml and unmarshall it in a structure for future use

Here we'll use Viper for Golang
It's a package that serves to read configuration from a file
The best format for configuration is a .yml file
*/
package config

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var ConfigLoaded Config

func Get(configFilePath string) {
	viper.SetConfigName("config")                // name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configFilePath)          //check in the provided file path specified with launching args
	viper.AddConfigPath(".")                     // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil { //Try to read the config file at the specified location and if failure returns error depending of the problem
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Sprintf("fatal error loading the config file: %s", err.Error()))
		} else {
			panic(fmt.Sprintf("fatal error reading the config file: %s", err.Error()))
		}
	}

	err := viper.Unmarshal(&ConfigLoaded) //trying to unload the data from file into the object
	if err != nil {
		panic(fmt.Sprintf("fatal error unmarshalling the config file: %s", err.Error()))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//Creating a mutex to lock the config during the live unmarshalling to prevent concurent access
		mutex := new(sync.Mutex)
		mutex.Lock()
		fmt.Println("Config changed, unmarshalling")
		err := viper.Unmarshal(&ConfigLoaded) //trying to unload the data from file into the object
		if err != nil {
			//unlocking in case of error to be sure that it's ok if there's a failover mechanism
			mutex.Unlock()
			panic(fmt.Sprintf("fatal error unmarshalling the config file: %s", err.Error()))
		}
		//unlocking anyway
		mutex.Unlock()
	})
}
