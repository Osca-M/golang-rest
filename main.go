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
	//if err != nil {
	//	return fmt.Fprintf(err)
	//}
	fmt.Println("Endpoint Hit: homePage")
}

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article


func allArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: allArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {

	Articles = []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}