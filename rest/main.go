package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json: "Title"`
	Desc    string `json: "Desc"`
	Content string `json: "Content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Somebody To Love", Desc: "Song by Queen", Content: "A Song to Love"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func allTestArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint Hit: POST All Articles Endpoint")
}

func allPutArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint Hit: PUT All articles here")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/postarticles", allTestArticles).Methods("POST")
	myRouter.HandleFunc("/putarticles", allPutArticles).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
