package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/viper"

	"github.com/kind84/cacoo/models"

	"github.com/julienschmidt/httprouter"
)

// GetUser returns user info
func GetUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	resp, err := http.Get(fmt.Sprintf("https://cacoo.com/api/v1/users/%s.json", os.Getenv("USER_ID")))
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var u models.User
	decoder.Decode(&u)
	fmt.Println(u)

	s := fmt.Sprintf("Hello %s", u.Nickname)
	w.Write([]byte(s))
}

// GetBasicInfo for the current user.
func GetBasicInfo(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	chUser := make(chan models.User)
	chDgrams := make(chan models.Diagrams)

	go func(chUser chan<- models.User) {
		fmt.Println("1")
		resp, err := http.Get(fmt.Sprintf("https://cacoo.com/api/v1/users/%s.json", viper.Get("user_id")))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"code": 500, "msg": "Error occurred while trying to retrieve user info."}`))
		}

		dCoder := json.NewDecoder(resp.Body)
		defer resp.Body.Close()

		var user models.User
		dCoder.Decode(&user)
		chUser <- user
		close(chUser)
	}(chUser)

	go func(chDgrams chan<- models.Diagrams) {
		fmt.Println("2")
		resp, err := http.Get(fmt.Sprintf("https://cacoo.com/api/v1/diagrams.json?apiKey=%s", viper.Get("api_key")))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"code": 500, "msg": "Error occurred while trying to retrieve diagrams info."}`))
		}

		dCoder := json.NewDecoder(resp.Body)
		defer resp.Body.Close()

		var dGrams models.Diagrams
		err = dCoder.Decode(&dGrams)
		if err != nil {
			panic(err)
		}
		chDgrams <- dGrams
		close(chDgrams)
	}(chDgrams)

	var r struct {
		FullName    string `json:"full_name"`
		TotDiagrams int    `json:"total_diagrams"`
	}

	uOk := false
	dOk := false
	var u models.User
	var d models.Diagrams
	for !uOk && !dOk {
		u, uOk = <-chUser
		r.FullName = u.Nickname
		d, dOk = <-chDgrams
		r.TotDiagrams = d.Count
	}
	fmt.Println("3")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code": 500, "msg": "Error occurred while trying to serialize response."}`))
	}
}
