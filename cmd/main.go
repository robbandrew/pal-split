package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homePage")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	var newArticle Article
	err := json.NewDecoder(r.Body).Decode(&newArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Articles = append(Articles, newArticle)
	json.NewEncoder(w).Encode(newArticle)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Article not found", http.StatusNotFound)

}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updatedArticle Article
	err := json.NewDecoder(r.Body).Decode(&updatedArticle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, article := range Articles {
		if article.Id == id {
			Articles[index] = updatedArticle
			json.NewEncoder(w).Encode(updatedArticle)
			return
		}
	}
	http.Error(w, "Article not found", http.StatusNotFound)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Endpoint Hit: returnSingleArticle")
	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	http.Error(w, "Article not found", http.StatusNotFound)
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/articles", returnAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", returnSingleArticle).Methods("GET")
	router.HandleFunc("/articles", createNewArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")

	fmt.Printf("Server started on %s", "http://localhost:10000")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {

	Articles = []Article{
		{Id: "1", Title: "First Article", Desc: "Description of first article", Content: "Content of first article"},
		{Id: "2", Title: "Second Article", Desc: "Description of second article", Content: "Content of second article"},
	}

	handleRequest()
}
