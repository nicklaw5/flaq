package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port       = "3000"
	streamName = "devstream"
	streamKey  = "DevStreamKey1"
)

func main() {
	flag.StringVar(&port, "port", port, "the port to serve on")
	flag.StringVar(&streamName, "name", streamName, "the name of the stream to serve")
	flag.StringVar(&streamKey, "key", streamKey, "the stream key used to authenticate when publishing")
	flag.Parse()

	http.HandleFunc("/streams/on-publish", postOnly(onPublish))
	http.HandleFunc("/streams/on-publish-done", postOnly(onPublishDone))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func postOnly(handler http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			handler(res, req)
			return
		}

		http.Error(res, "This endpopint only supports POST requests", http.StatusMethodNotAllowed)
	}
}

func onPublish(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	key := req.FormValue("name")

	// Check stream key exists. (Normally this would be matched against some sort of data store)
	if key != streamKey {
		// Invalid stream key, prevent stream from publishing
		http.Error(res, fmt.Sprintf("Invalid stream key \"%s\"", key), http.StatusUnauthorized)
		return
	}

	log.Printf("User has started streaming with stream key %s", streamKey)
	http.Redirect(res, req, fmt.Sprintf("rtmp://127.0.0.1/live/%s", streamName), http.StatusMovedPermanently)
}

func onPublishDone(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	streamKey := req.FormValue("name")

	log.Printf("User has stopped streaming with stream key %s", streamKey)
}
