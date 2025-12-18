package router

import (
	"github.com/VVUx21/mongoapi/controller"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}",controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movies/{id}",controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteall",controller.DeleteAllMovies).Methods("DELETE")
	return r
}
