package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type newreq struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	IsDone *bool  `json:"is_done"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var req Message
	json.NewDecoder(r.Body).Decode(&req)
	err := DB.Create(&req).Error
	if err != nil {
		log.Fatal("Failed to create Task in DB")
		return
	}
	fmt.Fprintln(w, "Task succesfully created")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)
	for _, task := range tasks {
		jsontask, err := json.Marshal(task)
		if err != nil {
			fmt.Fprintln(w, "Error with json.Marshal:", err)
			return
		}
		fmt.Fprintln(w, string(jsontask))
	}
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	var req newreq

	json.NewDecoder(r.Body).Decode(&req)
	result := DB.Model(&Message{}).Where("ID = ?", req.ID).Updates(&req)
	if result.Error != nil {
		fmt.Fprintln(w, "User not found:", result.Error)
		return
	} else {
		fmt.Fprintln(w, "Succesfully updated")
	}

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req newreq
	json.NewDecoder(r.Body).Decode(&req)

	result := DB.Where("ID = ?", req.ID).Delete(&Message{})
	if result.Error != nil {
		fmt.Fprintln(w, "Error with deleting user:", result.Error)
		return
	} else {
		fmt.Fprintf(w, "%v rows deleted", result.RowsAffected)
	}
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()

	router.HandleFunc("/api/create", PostHandler).Methods("POST")
	router.HandleFunc("/api/task", GetHandler).Methods("GET")
	router.HandleFunc("/api/update", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/delete", DeleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
