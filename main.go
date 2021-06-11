package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintf(w, "Welcome to the HomePage!")
	//if err != nil {
	//	return fmt.Fprintf(err)
	//}
	fmt.Println("Endpoint Hit: homePage")
}



