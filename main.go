package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)




func homePage(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

type Article struct {
	Id      string `json:"Id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article


func allArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: allArticles")
	json.NewEncoder(w).Encode(Articles)
}

func singleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/article/{id}", singleArticle)
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {

	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}