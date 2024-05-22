package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "go-clean-architecture/docs"
	. "go-clean-architecture/internal/domain"
	. "go-clean-architecture/internal/usecase"

	"github.com/gorilla/mux"
)

type ArticleSwag struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserName  string `json:"userName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ArticleController struct {
	useCase ArticleUseCase
}

func NewArticleController(router *mux.Router, us ArticleUseCase) {
	controller := ArticleController{
		useCase: us,
	}

	router.HandleFunc("/api/article", controller.CreateArticle).Methods("POST")
	router.HandleFunc("/api/article/{title}", controller.GetArticle).Methods("GET")
	router.HandleFunc("/api/articles", controller.GetArticles).Methods("GET")
	router.HandleFunc("/api/article/{title}", controller.DeleteArticle).Methods("DELETE")
}

// @Summary      Article 생성
// @Description  Article 을 생성합니다.
// @Tags         Article
// @ID			 create-article
// @Accept		 json
// @Produce      json
// @Param		 article body ArticleSwag true "새로운 Article 정보"
// @Success      201 {object} ArticleSwag
// @Router       /api/article [post]
func (c *ArticleController) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
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

// @Summary      단일 Article 조회
// @Description  단일 Article 을 조회합니다.
// @Tags         Article
// @ID			 get-article
// @Produce      json
// @Param        title path string true "조회할 Article 제목"
// @Success      200 {object} ArticleSwag
// @Router       /api/article/{title} [get]
func (c *ArticleController) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	article, _ := c.useCase.GetById(title)
	response, _ := json.Marshal(article)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}

// @Summary      Article 리스트 조회
// @Description  Article 리스트를 조회합니다.
// @Tags         Article
// @ID			 get-articles
// @Produce      json
// @Success      200 {array} ArticleSwag
// @Router       /api/articles [get]
func (c *ArticleController) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, _ := c.useCase.GetAll()
	response, _ := json.Marshal(articles)

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
// @Router       /api/article [delete]
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
