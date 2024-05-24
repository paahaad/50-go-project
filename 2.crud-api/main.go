package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Movie struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Rating int8 `json:"rating"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movie []Movie

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)
	for _, item := range  movie{
		if item.ID == param["id"]{
			json.NewEncoder(w).Encode(item)
			break
		}

	}
}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	var m Movie
	_ = json.NewDecoder(r.Body).Decode(&m)
	movie = append(movie, m)
	fmt.Println(movie)
	json.NewEncoder(w).Encode(m)
	
}

func updateMovies(){

}
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)

	for index, item := range movie{
		if item.ID == param["id"]{
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}
}



func main(){
	r := mux.NewRouter()

	movie = append(movie, Movie{ID: "1", Title: "Iron Man 1", Rating: 10, Director: &Director{Firstname: "Jon", Lastname: "Favreau"}})
	movie = append(movie, Movie{ID: "2", Title: "Iron Man 2", Rating: 9, Director: &Director{Firstname: "Jon", Lastname: "Favreau"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server is running at port 8080")
	if err := http.ListenAndServe(":8080", r); err!=nil{
		log.Fatal(err)
	}
}