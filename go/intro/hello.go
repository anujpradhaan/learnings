package main

import (
	"fmt"
	"log"

	"intro/mystring"
	"intro/packer"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Welcome to home page")
	fmt.Println("Writing to homepage")
}

func aboutPage(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Welcome to about page")
	fmt.Println("Writing to aboutpage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)

	// router.HandleFunc("/articles", allArticles)
	// router.HandleFunc("/article", createArticle).Methods(http.MethodPost)
	// router.HandleFunc("/article/{id}", getArticleByID).Methods(http.MethodGet)
	// router.HandleFunc("/article/{id}", updateArticle).Methods(http.MethodPut)
	// router.HandleFunc("/article/{id}", deleteArticle).Methods(http.MethodDelete)
	router.HandleFunc("/about", aboutPage)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func main() {
	fmt.Println("Hello")
	fmt.Println(packer.GetNumber())
	fmt.Println(mystring.GetString())
}
