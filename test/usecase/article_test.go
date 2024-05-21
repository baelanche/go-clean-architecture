package test

import (
	"errors"
	"go-clean-architecture/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockArticleRepository struct{}

func (m *MockArticleRepository) Create(article *domain.Article) (*domain.Article, error) {
	return nil, nil
}

func (m *MockArticleRepository) GetById(title string) (*domain.Article, error) {
	if title == "nonexistent" {
		return nil, errors.New("Article not found")
	}
	return &domain.Article{Title: title}, nil
}

func (m *MockArticleRepository) GetAll() ([]*domain.Article, error) {
	return []*domain.Article{}, nil
}

func (m *MockArticleRepository) Update(title string, article *domain.Article) (*domain.Article, error) {
	return nil, nil
}

func (m *MockArticleRepository) Delete(title string) error {
	return nil
}

func TestGetById(t *testing.T) {
	repo := &MockArticleRepository{}

	existingTitle := "existing"
	article, err := repo.GetById(existingTitle)
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, article, "Expected article not to be nil")
	assert.Equal(t, existingTitle, article.Title, "Expected article with title %q", existingTitle)

	nonexistentTitle := "nonexistent"
	article, err = repo.GetById(nonexistentTitle)
	assert.Error(t, err, "Expected error")
	assert.Nil(t, article, "Expected article to be nil")
}

func TestCreate(t *testing.T) {
	repo := &MockArticleRepository{}

	article := &domain.Article{Title: "new", Body: "body", UserName: "user"}
	createdArticle, err := repo.Create(article)
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, createdArticle, "Expected created article not to be nil")
}
