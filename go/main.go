package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type MeetingToken struct {
	Token string `json:"token"`
}
type tokenResponse struct {
	token string
}

var VIDEOSDK_API_KEY = ""
var VIDEOSDK_SECRET_KEY = ""
var VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live"

func getToken(w http.ResponseWriter, r *http.Request) {

	var permissions [2]string
	permissions[0] = "allow_join"
	permissions[1] = "allow_mod"

	var roles [2]string
	roles[0] = "CRAWLER"
	roles[1] = "PUBLISHER"

	atClaims := jwt.MapClaims{}
	atClaims["version"] = 2
	atClaims["apikey"] = VIDEOSDK_API_KEY
	atClaims["permissions"] = permissions
	atClaims["roles"] = roles
	atClaims["iat"] = time.Now().Unix()
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(VIDEOSDK_SECRET_KEY))

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// return token, nil
	var tokenString = token
	var res = `{"token":"` + tokenString + `" }`
	// var t tokenResponse

	in := []byte(res)
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	json.Marshal(res)

	// json.Unmarshal([]byte(res), &t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func createMeeting(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var str = []byte(string(b))

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var m MeetingToken
	json.Unmarshal([]byte(string(b)), &m)

	url := VIDEOSDK_API_ENDPOINT + "/api/meetings"
	method := "POST"

	// fmt.Print("\n", strings.NewReader(s))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(str))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", m.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(body))
}

func validateMeeting(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/validate-meeting/")

	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var str = []byte(string(b))

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var m MeetingToken
	json.Unmarshal([]byte(string(b)), &m)

	url := VIDEOSDK_API_ENDPOINT + "/api/meetings/" + id
	method := "POST"

	// fmt.Print("\n", strings.NewReader(s))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(str))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", m.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get-token", getToken).Methods(http.MethodGet)
	r.HandleFunc("/create-meeting", createMeeting).Methods(http.MethodPost)
	r.HandleFunc("/validate-meeting/{meetingId}", validateMeeting).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
