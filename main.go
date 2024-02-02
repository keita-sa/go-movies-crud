package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // declare a struct with a pointer type in the struct field
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie // creating slice to hold Movie data

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // setting Response header
	json.NewEncoder(w).Encode(movies)                  // encoding from Go data types to JSON
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // setting Response header
	params := mux.Vars(r)                              // getting path params using Gorilla/Mux
	for _, item := range movies {                      // Loop with "for _, v := range". _, means no need to set times (usually i).
		if item.ID == params["id"] { // if the ID of the item and the id of the params are the same
			json.NewEncoder(w).Encode(item) // encoding from Go data types to JSON
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // setting Response header
	var movie Movie                                    // new type defined in struct, variable movie of type Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)         // decoding the JSON data and store the result into the movie address of the Movie struct
	movie.ID = strconv.Itoa(rand.Intn(100))            // converting numbers to strings with strconv.Itoa()
	movies = append(movies, movie)                     // appending to the last element of the movies slice
	json.NewEncoder(w).Encode(movie)                   // encoding the JSON data
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // setting Response header
	params := mux.Vars(r)                              // getting query params

	// delete the movie with the i.d that you've sent
	// add a new movie - the movie that we send in the body of postman
	for index, item := range movies { // loop over the movies, range
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // adding the slices with append()(appending...after the second argument)
			var movie Movie                                      // New type defined in struct, variable movie of type Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)           // decoding JSON data and storing the value in the struct Movie, movie address.
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) // encoding JSON data
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // setting Response header
	params := mux.Vars(r)                              // getting path params with Gorilla/Mux
	for index, item := range movies {                  // loop over the movies, range

		if item.ID == params["id"] { //  if the ID of the item and the id of the params are the same
			movies = append(movies[:index], movies[index+1:]...) // appending slices with append()
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //encoding JSON data
}

func main() {
	r := mux.NewRouter() // create Multiplexers(Mux) and initialise routers with the NewRouter function

	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "454369", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
