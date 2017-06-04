package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	DevStream     = "devstream"
	DevStreamKey1 = "DevStreamKey1"
	HTTPPort      = 3000
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/streams/on-publish", onPublish).Methods("POST")
	r.HandleFunc("/streams/on-publish-done", onPublishDone).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", HTTPPort), r))
}

func onPublish(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	streamKey := req.FormValue("name")

	// Check stream key exists. (Normally this would be matched against some sort of data store)
	if streamKey != DevStreamKey1 {
		// Invalid stream key, prevent stream from publishing
		log.Printf("Invalid stream key %s", streamKey)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("User has started streaming with stream key %s", streamKey)
	http.Redirect(res, req, fmt.Sprintf("rtmp://127.0.0.1/live/%s", DevStream), 301)
}

func onPublishDone(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	streamKey := req.FormValue("name")

	log.Printf("User has stopped streaming with stream key %s", streamKey)
}
