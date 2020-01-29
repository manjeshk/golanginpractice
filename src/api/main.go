package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//Define array of articles
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Golang", Desc: "Go language", Content: "Concurrency"},
		Article{Title: "Java", Desc: "Java language", Content: "Multithreading"},
	}

	fmt.Println("Endpoint hit for all articles")
	json.NewEncoder(w).Encode(articles)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage endpoint hit")
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "postArticles endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

	//http.HandleFunc("/", homepage)
	//http.HandleFunc("/articles", allArticles)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
