package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//defined  book structure

type Book struct {
	Name    string `json:"Name"`
	Author  string `json:"Author"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var Books []Book

// after processing http request it will send http response in JSON format
func returnBooks(w http.ResponseWriter, r *http.Request) {
	books := []Book{
		Book{Name: "Book1", Author: "AuthorName1", Desc: "Description", Content: "Fiction"},
		Book{Name: "Book2", Author: "AuthorName2", Desc: "Description", Content: "Sci-Fi"},
		Book{Name: "Book3", Author: "AuthorName3", Desc: "Description", Content: "Horror"},
	}
	fmt.Println("Endpoint Hit: All books Endpoint")
	json.NewEncoder(w).Encode(books) //encoding our array into json string and then return as response

}

//process http request and send http response
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome World!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	//match the url with the defined function
	//handleRequests
	http.HandleFunc("/", helloWorld)

	http.HandleFunc("/books", returnBooks)
	//fatal is used to print error messages
	log.Fatal(http.ListenAndServe(":8082", nil))
}
