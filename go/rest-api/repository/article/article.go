package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}

// var Articles []Article = []Article{
// 	Article{ID: "1", Title: "Hello", Description: "Hello Desc", Content: "Hello Content"},
// 	Article{ID: "2", Title: "Bye", Description: "Bye Desc", Content: "Bye Content"},
// }

func allArticles(responseWriter http.ResponseWriter, request *http.Request) {
	json.NewEncoder(responseWriter).Encode(Articles)
}

func getArticleByID(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(responseWriter).Encode(article)
		}
	}
}

func createArticle(responseWriter http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)

	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(responseWriter).Encode(Articles)
}

func updateArticle(responseWriter http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	vars := mux.Vars(request)
	id := vars["id"]

	var articleToUpdate Article
	json.Unmarshal(reqBody, &articleToUpdate)

	for index, article := range Articles {
		if article.ID == id {
			articleToUpdate.ID = id
			Articles[index] = articleToUpdate
			// Articles[index].Content = articleToUpdate.Content
			// Articles[index].Description = articleToUpdate.Description
			// Articles[index].Title = articleToUpdate.Title

			json.NewEncoder(responseWriter).Encode(Articles)
		}
	}

}

func deleteArticle(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println("deleting article with id = " + id)

	for index, article := range Articles {
		if article.ID == id {
			fmt.Println("Found id = " + id)
			Articles = append(Articles[:index], Articles[index+1:]...)
			json.NewEncoder(responseWriter).Encode(Articles)
		}
	}
}
