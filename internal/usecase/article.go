package usecase

import (
	"go-clean-architecture/internal/model"
)

type ArticleService struct {
	article model.ArticleRepository
}

func NewArticleService(article model.ArticleRepository) model.ArticleUseCase {
	return &ArticleService{
		article,
	}
}

func (service *ArticleService) Create(article *model.Article) (*model.Article, error) {
	return service.article.Create(article)
}

func (service *ArticleService) GetById(title string) (*model.Article, error) {
	return service.article.GetById(title)
}
func (service *ArticleService) GetAll() ([]*model.Article, error) {
	return service.article.GetAll()
}

func (service *ArticleService) Update(title string, article *model.Article) (*model.Article, error) {
	return service.article.Update(title, article)
}

func (service *ArticleService) Delete(title string) error {
	return service.article.Delete(title)
}
