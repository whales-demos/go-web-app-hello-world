// Package main : a simple web app
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
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
	if httpPort == "" {
		httpPort = "8080"
	}

	var message = os.Getenv("MESSAGE")
	if message == "" {
		message = "this is a message"
	}

	log.Println("🚀 starting web server on port: " + httpPort)
	log.Println("📝 message: " + message)

	var authenticationStatus = ""
	var dbPasswordFile = os.Getenv("DB_PASSWORD_FILE") // 🤭 I know it is not a good practice to store passwords in files

	log.Println("dbPasswordFile", dbPasswordFile)

	if dbPasswordFile == "" {

		authenticationStatus = "😡 you are not authenticated"

		log.Println("dbPasswordFile empty", authenticationStatus)

	} else {
		// read content of file, the path of the file is in dbPasswordFile
		// if the content is "password" then set authenticationStatus to "you are authenticated"
		// else set authenticationStatus to "you are not authenticated"
		content, err := os.ReadFile(dbPasswordFile)
		if err != nil {
			authenticationStatus = "😡 you are not authenticated: " + err.Error()

			log.Println("dbPasswordFile error", authenticationStatus)

		}

		log.Println("📝 dbPasswordFile content", content, string(content))

		if strings.Replace(string(content), "\n", "", -1) == "password" {
			authenticationStatus = "🙂 you are authenticated"
		} else {
			authenticationStatus = "😡 you are not authenticated"

			log.Println("false content", authenticationStatus)

		}
	}

	mux := http.NewServeMux()

	fileServerHtml := http.FileServer(http.Dir("public"))
	mux.Handle("/", fileServerHtml)

	mux.HandleFunc("/variables", func(response http.ResponseWriter, request *http.Request) {

		variables := map[string]interface{}{
			"message":              message,
			"authenticationStatus": authenticationStatus,
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
