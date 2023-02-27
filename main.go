package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"  //new id for new movie
	"net/http" //create server for golang
	"strconv" //id in integer and to convert into string
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`
}

type Director struct{
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){  //passing a * of the request from postman to this fn...
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) //pass the complete movie slice...

}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movie[index + 1:]...) //the movie that i want to delete will be caught up by the index and replaced by remaining ...
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies{  //48.10 yt
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewEncoder(r.Body).Encode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	//set json content type
	w.Header().Set("Content-Type", "application/json")

	//params
	params := mux.Vars(r)
	//loop over the movies, range 
	//delete the movie id that sent
	//add a new movie- that in the body of postman

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_= json.NewEncoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

//drivers code...
func main() {
	r := mux.NewRouter()
	
	movies = append(movies, Movie
		{ ID:1, Isbn: "95533", Title: "Movie One", Director: &Director{FirstName:"Bruce", LastName:"Wayne"}})
	movies = append(movies, Movie
		{ ID:2, Isbn: "95534", Title: "Movie Two", Director: &Director{FirstName:"Tom", LastName:"Latham"}})

	// five functions which we can do CRUD...
	r.HandleFunc("/movies"     , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies"     , createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenandServe(":8080", r)) //to start a server we use this.
}