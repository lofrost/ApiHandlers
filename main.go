package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Msg string `json:"message"`
}

var task string

func PostHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)

	var req Message

	d.Decode(&req)
	err := DB.Create(&req).Error
	if err != nil {
		log.Fatal("Failed to create Task in DB")
	}
	fmt.Fprintln(w, "Task succesfully created")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)
	//fmt.Fprintln(w, tasks)
	for _, task := range tasks {
		fmt.Fprintf(w, "Task - %v\nIsDone - %v\nCreated at - %v\n\n", task.Task, task.IsDone, task.CreatedAt)
	}
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()

	router.HandleFunc("/api/update", PostHandler).Methods("POST")
	router.HandleFunc("/api/task", GetHandler).Methods("GET")
	http.ListenAndServe(":8080", router)

}
