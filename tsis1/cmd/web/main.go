package main

import (
	"fmt"
	"net/http"

	"github.com/Nursat22B030486/tsis1/internal/movie"

	"encoding/json"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies", GetList).Methods("GET")
	router.HandleFunc("/specific-movie", GetInfoMovie).Methods("GET")
	router.HandleFunc("/actor-movies", GetMovieByActor).Methods("GET")
	// router.HandleFunc("/add-new-movie", AddNewMovie).Methods("PUT")
	router.HandleFunc("/about", About).Methods("Get")

	http.Handle("/", router)

	http.ListenAndServe(":8088", router)
}

func About(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,
		"Welcome to MOVIE Web Aplication!!!\n Web Application created by Nursat \n In this web application you can see various type of movies.")
}

func GetList(w http.ResponseWriter, r *http.Request) {

	var response movie.Movies

	films := movie.Database()

	response.Movies = films

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func GetInfoMovie(w http.ResponseWriter, r *http.Request) {
	var response movie.Movie

	info := movie.GetInfo("Ghosted")

	response = info

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func GetMovieByActor(w http.ResponseWriter, r *http.Request) {
	var response []movie.Movie

	movies := movie.GetMovies("Ryan Reynolds")

	response = movies

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
