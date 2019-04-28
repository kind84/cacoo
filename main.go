package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/spf13/viper"

	"github.com/kind84/cacoo/handlers"
	"github.com/kind84/cacoo/repo"
	"github.com/kind84/cacoo/stower"
)

func init() {
	fmt.Println("Setting up configuration...")
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
	// Repository singleton.
	Repo := repo.NewRedisRepo("127.0.0.1:6379")
	Stower := stower.NewStower(Repo)

	mux := httprouter.New()

	mux.GET("/user", handlers.GetUser)
	mux.GET("/api/basicInfo", handlers.GetBasicInfo(Repo, Stower))
	mux.GET("/api/folder/:id", handlers.GetFolder(Repo))

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", mux)
	// http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
