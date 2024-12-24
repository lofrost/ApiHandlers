package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Message string `json:"message"`
}

var task string

func PostHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)

	var req requestBody

	d.Decode(&req)

	task = req.Message
	fmt.Fprintln(w, "Task succesfully updated")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello, %s", task)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/update", PostHandler).Methods("POST")
	router.HandleFunc("/api/task", GetHandler).Methods("GET")
	http.ListenAndServe(":8080", router)

}
