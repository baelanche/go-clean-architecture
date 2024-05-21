package usecase

import (
	. "go-clean-architecture/internal/domain"
	. "go-clean-architecture/internal/repository/mysql"
)

type ArticleUseCase interface {
	Create(article *Article) (*Article, error)
	GetById(title string) (res *Article, err error)
	GetAll() (res []*Article, err error)
	Update(title string, article *Article) (*Article, error)
	Delete(title string) error
}

type articleUseCase struct {
	articleRepo ArticleRepository
}

func NewArticleUseCase(repo ArticleRepository) ArticleUseCase {
	return &articleUseCase{
		articleRepo: repo,
	}
}

func (repo *articleUseCase) Create(article *Article) (*Article, error) {
	return repo.articleRepo.Create(article)
}

func (repo *articleUseCase) GetById(title string) (*Article, error) {
	res, err := repo.articleRepo.GetById(title)
	return res, err
}
func (repo *articleUseCase) GetAll() ([]*Article, error) {
	return repo.articleRepo.GetAll()
}

func (repo *articleUseCase) Update(title string, article *Article) (*Article, error) {
	return repo.articleRepo.Update(title, article)
}

func (repo *articleUseCase) Delete(title string) error {
	return repo.articleRepo.Delete(title)
}
