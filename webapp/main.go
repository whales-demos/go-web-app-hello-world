// Package main : a simple web app
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

/*
GetBytesBody returns the body of an HTTP request as a []byte.
  - It takes a pointer to an http.Request as a parameter.
  - It returns a []byte.

func GetBytesBody(request *http.Request) []byte {
	body := make([]byte, request.ContentLength)
	request.Body.Read(body)
	return body
}
*/


func main() {

	var httpPort = os.Getenv("HTTP_PORT")
	if httpPort == "" { httpPort = "8080" }

	var message = os.Getenv("MESSAGE")
	if message == "" { message = "this is a message" }
	
	mux := http.NewServeMux()

	fileServerHtml := http.FileServer(http.Dir("public"))
	mux.Handle("/", fileServerHtml)

	mux.HandleFunc("/variables", func(response http.ResponseWriter, request *http.Request) {

		variables := map[string]interface{}{
			"message": message,
		}

		jsonString, err := json.Marshal(variables)

		if err != nil {
			response.WriteHeader(http.StatusNoContent)
		}

		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		response.WriteHeader(http.StatusOK)
		response.Write(jsonString)
	})

	var errListening error
	log.Println("🌍 http server is listening on: " + httpPort)
	errListening = http.ListenAndServe(":"+httpPort, mux)

	log.Fatal(errListening)
}
