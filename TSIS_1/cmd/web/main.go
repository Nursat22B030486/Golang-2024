package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type Movies struct {
	Movies []Movie `json:"movies"`
}

type Movie struct {
	Title    string  `json:"title"`
	Released int     `json:"realized"`
	Genre    string  `json:"genre"`
	Casts    []Actor `json:"casts"`
	Duration int     `json:"duration"`
	Country  string  `json:"country"`
}

type Actor struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
}

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

	var response Movies

	films := database()

	response.Movies = films

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func database() []Movie {
	var movies []Movie
	var movie Movie

	movie.Title = "Red Notice"
	movie.Released = 2021
	movie.Genre = "Comedy, Action, Thriller, Crime"
	movie.Casts = []Actor{
		{FirstName: "Ryan", LastName: "Reynolds", BirthDate: "23-10-1976"},
		{FirstName: "Gal", LastName: "Gadot", BirthDate: "30-04-1985"},
		{FirstName: "Dwayne ", LastName: "Johnson", BirthDate: "02-05-1972"},
	}
	movie.Duration = 116
	movie.Country = "United States of America"
	movies = append(movies, movie)

	movie.Title = "Bullet Train"
	movie.Released = 2022
	movie.Genre = " Action, Crime, Comedy"
	movie.Casts = []Actor{
		{FirstName: "Brad", LastName: "Pitt", BirthDate: "18-12-1963"},
		{FirstName: "Aaron", LastName: "Taylor-Johnson", BirthDate: "13-06-1990"},
		{FirstName: "Joey ", LastName: "King", BirthDate: "30-07-1999"},
	}
	movie.Duration = 127
	movie.Country = "Japan, United States of America"
	movies = append(movies, movie)

	movie.Title = "Ghosted"
	movie.Released = 2023
	movie.Genre = " Action, Comedy, Romance"
	movie.Casts = []Actor{
		{FirstName: "Chir", LastName: "Evans", BirthDate: "13-06-1981"},
		{FirstName: "Ana", LastName: "de Armas", BirthDate: "30-04-1988"},
		{FirstName: "Ryan", LastName: "Reynolds", BirthDate: "23-10-1976"},
	}
	movie.Duration = 116
	movie.Country = "United States of America"
	movies = append(movies, movie)

	movie.Title = "Fast and Furious 5"
	movie.Released = 2011
	movie.Genre = "Action, Crime, Thriller"
	movie.Casts = []Actor{
		{FirstName: "Vin", LastName: "Diesel", BirthDate: "18-07-1967"},
		{FirstName: "Paul", LastName: "Walker", BirthDate: "12-09-1973"},
		{FirstName: "Dwayne", LastName: "Johnson", BirthDate: "02-05-1972"},
		{FirstName: "Gal", LastName: "Gadot", BirthDate: "30-04-1985"},
	}
	movie.Duration = 130
	movie.Country = "United States of America"
	movies = append(movies, movie)

	movie.Title = "Avengers: Endgame"
	movie.Released = 2019
	movie.Genre = "Action, Adventure, Drama"
	movie.Casts = []Actor{
		{FirstName: "Robert", LastName: "Downey Jr.", BirthDate: "04-04-1965"},
		{FirstName: "Chris", LastName: "Evans", BirthDate: "13-06-1981"},
		{FirstName: "Scarlett", LastName: "Johansson", BirthDate: "22-11-1984"},
	}
	movie.Duration = 181
	movie.Country = "United States of America"
	movies = append(movies, movie)

	return movies
}

func GetInfoMovie(w http.ResponseWriter, r *http.Request) {
	var response Movie

	info := getInfo("Ghosted")

	response = info

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func getInfo(title string) Movie {
	var response Movie
	movies := database()

	for _, movie := range movies {
		if movie.Title == title {
			response = movie
		}
	}
	return response
}

func GetMovieByActor(w http.ResponseWriter, r *http.Request) {
	var response []Movie

	movies := getMovies("Ryan Reynolds")

	response = movies

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func getMovies(name string) []Movie {
	var response []Movie

	movies := database()

	for _, movie := range movies {
		for _, actor := range movie.Casts {
			fullname := actor.FirstName + " " + actor.LastName
			fmt.Println(fullname)
			if name == fullname {
				response = append(response, movie)
			}
		}
	}
	return response
}
