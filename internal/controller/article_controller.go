package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-clean-architecture/internal/domain"

	"github.com/gorilla/mux"
)

type ArticleController struct {
	useCase domain.ArticleUseCase
}

func NewArticleController(router *mux.Router, us domain.ArticleUseCase) {
	controller := ArticleController{
		useCase: us,
	}

	router.HandleFunc("/article", controller.CreateArticle).Methods("POST")
	router.HandleFunc("/article/{title}", controller.GetArticle).Methods("GET")
	router.HandleFunc("/articles", controller.GetArticles).Methods("GET")
	router.HandleFunc("/article/{title}", controller.DeleteArticle).Methods("DELETE")
}

func (c *ArticleController) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article domain.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	done, err := c.useCase.Create(&article)
	if err != nil {
		http.Error(w, "Failed to create Article", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(done)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

func (c *ArticleController) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	article, _ := c.useCase.GetById(title)
	response, _ := json.Marshal(article)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

func (c *ArticleController) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, _ := c.useCase.GetAll()
	response, _ := json.Marshal(articles)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

func (c *ArticleController) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	err := c.useCase.Delete(title)
	if err != nil {
		http.Error(w, "Article does not exist", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Delete was successful")
}
