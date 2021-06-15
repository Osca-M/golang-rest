# golang-rest

A simple Golang REST API that encompasses CRUD methodologies for an Article struct that is stored in memory.
```
type Article struct {
Id      string `json:"Id"`
Title   string `json:"title"`
Desc    string `json:"desc"`
Content string `json:"content"`
}
```
To get these functionalities, clone the repository and
run ```go run main.go```

I hope this illustrates how to write a basic CRUD REST API using Golang and gorilla/mux.

Credits to this guide by [TutorialEdge.net](https://tutorialedge.net/golang/creating-restful-api-with-golang/)