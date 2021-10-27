package router

import (
	"github.com/pbtrad/golang-react-todo/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks/{id", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask", middleware.undoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAlltasks", middleware.deleteAlltasks).Methods("DELETE", "OPTIONS")
	return router
}
