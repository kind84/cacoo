package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/kind84/cacoo/handlers"

	"github.com/julienschmidt/httprouter"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("cacoo")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func main() {
	mux := httprouter.New()
	mux.GET("/user", handlers.GetUser)
	mux.GET("/api/basicInfo", handlers.GetBasicInfo)
	// http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	http.ListenAndServe(":8080", mux)
}
