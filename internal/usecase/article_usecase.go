package usecase

import (
	"go-clean-architecture/internal/domain"
)

type articleUseCase struct {
	articleRepo domain.ArticleRepository
}

func NewArticleUseCase(repo domain.ArticleRepository) domain.ArticleUseCase {
	return &articleUseCase{
		articleRepo: repo,
	}
}

func (repo *articleUseCase) Create(article *domain.Article) (*domain.Article, error) {
	return repo.articleRepo.Create(article)
}

func (repo *articleUseCase) GetById(title string) (*domain.Article, error) {
	res, err := repo.articleRepo.GetById(title)
	return res, err
}
func (repo *articleUseCase) GetAll() ([]*domain.Article, error) {
	return repo.articleRepo.GetAll()
}

func (repo *articleUseCase) Update(title string, article *domain.Article) (*domain.Article, error) {
	return repo.articleRepo.Update(title, article)
}

func (repo *articleUseCase) Delete(title string) error {
	return repo.articleRepo.Delete(title)
}
