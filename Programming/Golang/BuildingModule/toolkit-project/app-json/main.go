package main

import (
	"log"
	"net/http"
	"toolkit"
)

type RequestPayload struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ResponsePayload struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/receive-post", receivePost)
	mux.HandleFunc("/remote-service", remoteService)
	mux.HandleFunc("/simulated-service", simulatedService)

	log.Println("Starting the server...")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func receivePost(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools
	err := t.ReadJson(w, r, &requestPayload)
	if err != nil {
		t.ErrorJson(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message: "The handler works correctly",
	}
	err = t.WriteJson(w, http.StatusAccepted, responsePayload)
	if err != nil {
		log.Println(err)
	}
}

func remoteService(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools
	err := t.ReadJson(w, r, &requestPayload)
	if err != nil {
		t.ErrorJson(w, err)
		return
	}

	_, statusCode, err := t.PushJsonToRemote("http://localhost:8081/simulated-service", requestPayload)
	if err != nil {
		t.ErrorJson(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message:    "The remote service handler works correctly",
		StatusCode: statusCode,
	}
	err = t.WriteJson(w, http.StatusAccepted, responsePayload)
	if err != nil {
		log.Println(err)
	}
}

func simulatedService(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	payload := ResponsePayload{
		Message: "ok",
	}
	_ = t.WriteJson(w, http.StatusOK, payload)
}
