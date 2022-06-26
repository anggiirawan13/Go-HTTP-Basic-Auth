package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("start main.go")
}

func main() {
	http.HandleFunc("/students", ActionStudents)

	server := new(http.Server)
	server.Addr = ":8080"

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ActionStudents(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if !Auth(w, r) { return }
	if !AllowOnlyMethodGET(w, r) { return }

	idString := r.URL.Query().Get("id")
	if strings.TrimSpace(idString) != "" {
		id, err := strconv.Atoi(idString)

		if err != nil {
			fmt.Println("func ActionStudents")

			response, _ := json.Marshal(DefaultResponse{
				Success:    false,
				Messages:   "failed boss",
				Data:       nil,
				StatusCode: http.StatusBadRequest,
			})

			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(response)

			return
		}

		if id > 0 {
			OutputJSON(w, SelectStudents(int64(id)))
			return
		}
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, any interface{})  {
	resp := DefaultResponse{
		Success:    true,
		Messages:   "successed boss",
		Data:       any,
		StatusCode: http.StatusOK,
	}

	response, _ := json.Marshal(resp)

	_, err := w.Write(response)
	if err != nil {
		fmt.Println("func OutputJSON err")

		response, _ := json.Marshal(DefaultResponse{
			Success:    false,
			Messages:   "failed boss",
			Data:       nil,
			StatusCode: http.StatusBadRequest,
		})

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(response)

		return
	}
}