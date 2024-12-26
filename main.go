package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type newreq struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	IsDone *bool  `json:"is_done"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var req Message
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")
	err := DB.Create(&req).Error
	if err != nil {
		resp := Response{
			Status:  "error",
			Message: "Failed to create Task in DB",
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	task, err := json.Marshal(req)
	if err != nil {
		log.Fatal("Error with json.Marshal")
		return
	}
	w.Write(task)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	DB.Find(&tasks)
	jsontasks, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal("Error with json.Marshal")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsontasks)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	var req newreq
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&req)
	result := DB.Model(&Message{}).Where("ID = ?", req.ID).Updates(&req)
	if result.Error != nil {
		resp := Response{
			Status:  "error",
			Message: "DB error:" + result.Error.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	} else if result.RowsAffected == 0 {
		resp := Response{
			Status:  "error",
			Message: "Task not found",
		}
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		var UpdatedMessage Message
		DB.Model(&Message{}).Where("ID = ?", req.ID).First(&UpdatedMessage)
		json.NewEncoder(w).Encode(UpdatedMessage)
	}

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req newreq
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Set("Content-Type", "application/json")
	result := DB.Where("ID = ?", req.ID).Delete(&Message{})
	if result.Error != nil {
		resp := Response{
			Status:  "error",
			Message: "DB error:" + result.Error.Error(),
		}
		json.NewEncoder(w).Encode(resp)
		return
	} else if result.RowsAffected == 0 {
		resp := Response{
			Status:  "error",
			Message: "Task not found",
		}
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		resp := Response{
			Status:  "success",
			Message: "Task deleted",
		}
		json.NewEncoder(w).Encode(resp)
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
