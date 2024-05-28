package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "go-clean-architecture/docs"
	"go-clean-architecture/internal/model"

	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	article model.ArticleUseCase
}

func NewArticleHandler(router *mux.Router, article model.ArticleUseCase) {
	handler := ArticleHandler{
		article,
	}

	router.HandleFunc("/api/article", handler.CreateArticle).Methods("POST")
	router.HandleFunc("/api/article/{title}", handler.GetArticle).Methods("GET")
	router.HandleFunc("/api/articles", handler.GetArticles).Methods("GET")
	router.HandleFunc("/api/article/{title}", handler.UpdateArticle).Methods("PUT")
	router.HandleFunc("/api/article/{title}", handler.DeleteArticle).Methods("DELETE")
}

// @Summary      Article 생성
// @Description  Article 을 생성합니다.
// @Tags         Article
// @ID			 create-article
// @Accept		 json
// @Produce      json
// @Param		 article body model.Article true "새로운 Article 정보"
// @Success      201 {object} model.Article
// @Router       /api/article [post]
func (handler *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	newArticle, err := handler.article.Create(&article)
	if err != nil {
		http.Error(w, "Failed to create Article", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(newArticle)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

// @Summary      단일 Article 조회
// @Description  단일 Article 을 조회합니다.
// @Tags         Article
// @ID			 get-article
// @Produce      json
// @Param        title path string true "조회할 Article 제목"
// @Success      200 {object} model.Article
// @Router       /api/article/{title} [get]
func (handler *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	article, _ := handler.article.GetById(title)
	response, _ := json.Marshal(article)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

// @Summary      Article 리스트 조회
// @Description  Article 리스트를 조회합니다.
// @Tags         Article
// @ID			 get-articles
// @Produce      json
// @Success      200 {array} model.Article
// @Router       /api/articles [get]
func (handler *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, _ := handler.article.GetAll()
	response, _ := json.Marshal(articles)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

// @Summary      Article 수정
// @Description  Article 을 수정합니다.
// @Tags         Article
// @ID			 update-article
// @Accept		 json
// @Produce      json
// @Param		 title path string true "수정할 Article 제목"
// @Param		 article body model.Article true "수정할 Article 정보"
// @Success      200 {object} string
// @Router       /api/article/{title} [put]
func (handler *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	_, err := handler.article.GetById(title)
	if err != nil {
		http.Error(w, "Article does not exist", http.StatusBadRequest)
		return
	}

	var article model.Article
	err = json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	updatedArticle, err := handler.article.Update(title, &article)
	if err != nil {
		http.Error(w, "Failed to update Article", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(updatedArticle)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

// @Summary      Article 제거
// @Description  Article 을 제거합니다.
// @Tags         Article
// @ID			 delete-article
// @Accept		 json
// @Produce      json
// @Param		 title path string true "삭제할 Article 제목"
// @Success      200 {object} string
// @Router       /api/article/{title} [delete]
func (handler *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	err := handler.article.Delete(title)
	if err != nil {
		http.Error(w, "Article does not exist", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Delete was successful")
}
