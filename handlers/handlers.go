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

type cacooReq struct {
	index  int
	url    string
	errMsg string
}

type cacooResp struct {
	index int
	url   string
	resp  *http.Response
}

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
	// chUser := make(chan models.User)
	// chDgrams := make(chan models.Diagrams)
	fmt.Println("Retrieving basic info...")
	requests := map[int]cacooReq{
		0: {
			0,
			fmt.Sprintf("https://cacoo.com/api/v1/users/%s.json", viper.Get("user_id")),
			`{"code": 500, "msg": "Error occurred while trying to retrieve user info."}`,
		},
		1: {
			1,
			fmt.Sprintf("https://cacoo.com/api/v1/diagrams.json?apiKey=%s", viper.Get("api_key")),
			`{"code": 500, "msg": "Error occurred while trying to retrieve diagrams info."}`,
		},
		2: {
			2,
			fmt.Sprintf("https://cacoo.com/api/v1/folders.json?apiKey=%s", viper.Get("api_key")),
			`{"code": 500, "msg": "Error occurred while trying to retrieve folders info."}`,
		},
	}

	chResp := make(chan cacooResp, len(requests))

	for _, r := range requests {
		go cacooCall(r, chResp, w)
	}

	var r struct {
		FullName    string `json:"full_name"`
		TotDiagrams int    `json:"total_diagrams"`
		Folders     []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
	}

	resCount := 0
	for cResp := range chResp {
		fmt.Printf("Got response from %s\n", cResp.url)
		dCoder := json.NewDecoder(cResp.resp.Body)
		defer cResp.resp.Body.Close()
		switch cResp.index {
		case 0:
			var mUsr models.User
			dCoder.Decode(&mUsr)
			fmt.Printf("Decoded user %v\n", mUsr)
			r.FullName = mUsr.Nickname
			resCount++
		case 1:
			var mDgr models.Diagrams
			dCoder.Decode(&mDgr)
			fmt.Printf("Decoded diagrams %v\n", mDgr)
			r.TotDiagrams = mDgr.Count
			resCount++
		case 2:
			var mFold models.Folders
			dCoder.Decode(&mFold)
			fmt.Printf("Decoded folders %v\n", mFold)
			for _, f := range mFold.Result {
				r.Folders = append(r.Folders, struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				}{ID: f.FolderID, Name: f.FolderName})
			}
			resCount++
		}
		if resCount == len(requests) {
			break
		}
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code": 500, "msg": "Error occurred while trying to serialize response."}`))
	}
}

func cacooCall(req cacooReq, chResp chan<- cacooResp, w http.ResponseWriter) {
	fmt.Printf("Calling %s ...\n", req.url)
	resp, err := http.Get(req.url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(req.errMsg))
	}

	chResp <- cacooResp{
		index: req.index,
		url:   req.url,
		resp:  resp,
	}
}
