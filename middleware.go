package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Request header
// Authorization: Basic c29tZXVzZXJuYW1lOnNvbWVwYXNzd29yZA==
const (
	USERNAME = "someusername"
	PASSWORD = "somepassword"
)

func init() {
	fmt.Println("start middleware.go")
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		fmt.Println("func Auth")

		response, _ := json.Marshal(DefaultResponse{
			Success:    false,
			Messages:   "failed boss",
			Data:       nil,
			StatusCode: http.StatusBadRequest,
		})

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(response)

		return false
	}

	isValid := (USERNAME == username) && (PASSWORD == password)
	if !isValid {
		fmt.Println("func Auth isValid")

		response, _ := json.Marshal(DefaultResponse{
			Success:    false,
			Messages:   "failed boss",
			Data:       nil,
			StatusCode: http.StatusBadRequest,
		})

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(response)

		return false
	}

	return true
}

func AllowOnlyMethodGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		fmt.Println("func AllowOnlyMethodGET")

		response, _ := json.Marshal(DefaultResponse{
			Success:    false,
			Messages:   "failed boss",
			Data:       nil,
			StatusCode: http.StatusBadRequest,
		})

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(response)

		return false
	}

	return true
}
