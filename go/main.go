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
var VIDEOSDK_API_ENDPOINT = "https://api.videosdk.live/v2"

func getToken(w http.ResponseWriter, r *http.Request) {

	var permissions [2]string
	permissions[0] = "allow_join"
	permissions[1] = "allow_mod"
	var roles [2]string
	roles[0] = "crawler"
	roles[1] = "rtc"

	atClaims := jwt.MapClaims{}
	atClaims["apikey"] = VIDEOSDK_API_KEY
	atClaims["permissions"] = permissions
	atClaims["iat"] = time.Now().Unix()
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// atClaims["version"] = 2  //OPTIONAL  // For accessing the v2 API, For passing roomId, participantId or roles parameters in payload it is required to pass.
    // atClaims["roomId"] = "2kyv-gzay-64pg" //OPTIONAL // To provide customised access control, you can make the token applicable to a particular room by including the roomId in the payload.
    // atClaims["participantId"]= "lxvdplwt" //OPTIONAL  // You can include the participantId in the payload to limit the token's access to a particular participant.
    // atClaims["roles"] = roles //OPTIONAL // crawler role is only for accessing v2 API, you can not use this token for running the Meeting/Room. rtc is only allow for running the Meeting / Room, you can not use server-side APIs.

    //OPTIONALLY add the version, roles, roomId, and peerId if you wish to use this token for joining the meeeting
    //with a particular roomId or participantId
	if roomId, peerId := r.URL.Query().Get("roomId"), r.URL.Query().Get("peerId"); roomId != "" || peerId != "" {
		atClaims["version"] = 2
		atClaims["roles"] = []string{"rtc"}
		if roomId != "" {
			atClaims["roomId"] = roomId
		}
		if peerId != "" {
			atClaims["participantId"] = peerId
		}
	}

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

	url := VIDEOSDK_API_ENDPOINT + "/rooms"
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

    var m MeetingToken
    if err := json.Unmarshal(b, &m); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    url := VIDEOSDK_API_ENDPOINT + "/rooms/validate/" + id
    method := "GET"

    client := &http.Client{}
    req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
    if err != nil {
        fmt.Println(err)
        http.Error(w, err.Error(), 500)
        return
    }

    req.Header.Add("Authorization", m.Token)
    req.Header.Add("Content-Type", "application/json")

    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        http.Error(w, err.Error(), 500)
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        http.Error(w, err.Error(), 500)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(body)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get-token", getToken).Methods(http.MethodGet)
	r.HandleFunc("/create-meeting", createMeeting).Methods(http.MethodPost)
	r.HandleFunc("/validate-meeting/{meetingId}", validateMeeting).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
