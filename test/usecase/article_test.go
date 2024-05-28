package test

import (
	"errors"
	"go-clean-architecture/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockArticleMySqlRepository struct {
	articles []model.Article
}

func (m *MockArticleMySqlRepository) Create(article *model.Article) (*model.Article, error) {
	m.articles = append(m.articles, *article)
	return article, nil
}

func (m *MockArticleMySqlRepository) GetById(title string) (*model.Article, error) {
	for _, article := range m.articles {
		if article.Title == title {
			return &article, nil
		}
	}
	return nil, errors.New("Article not found")
}

func (m *MockArticleMySqlRepository) GetAll() ([]model.Article, error) {
	return m.articles, nil
}

func (m *MockArticleMySqlRepository) Update(title string, updatedArticle *model.Article) (*model.Article, error) {
	for i, article := range m.articles {
		if article.Title == title {
			m.articles[i] = *updatedArticle
			return updatedArticle, nil
		}
	}

	return nil, errors.New("Article not found")
}

func (m *MockArticleMySqlRepository) Delete(title string) error {
	for i, article := range m.articles {
		if article.Title == title {
			m.articles = append(m.articles[:i], m.articles[i+1:]...)
			return nil
		}
	}
	return errors.New("Article not found")
}

func TestArticle(t *testing.T) {
	repo := &MockArticleMySqlRepository{}

	t.Run("Create", func(t *testing.T) {
		article := &model.Article{Title: "new", Body: "body", UserName: "user"}
		createdArticle, err := repo.Create(article)
		assert.NoError(t, err, "Expected no error")
		assert.NotNil(t, createdArticle, "Expected created article not to be nil")
	})

	t.Run("GetById", func(t *testing.T) {
		title := "new"
		article, err := repo.GetById(title)
		assert.NoError(t, err, "Expected no error")
		assert.NotNil(t, article, "Expected article not to be nil")
		assert.Equal(t, article.Title, title, "Expected article with title %q", title)
	})

	t.Run("GetAll", func(t *testing.T) {
		articles, err := repo.GetAll()
		assert.NoError(t, err, "Expected no error")
		assert.NotNil(t, articles, "Expected articles not to be nil")
		assert.Equal(t, len(articles), 1, "Expected length of articles is one")
	})

	t.Run("Update", func(t *testing.T) {
		article := &model.Article{Title: "update", Body: "body", UserName: "user"}
		updatedArticle, err := repo.Update("new", article)
		assert.NoError(t, err, "Expected no error")
		assert.NotNil(t, updatedArticle, "Expected article not to be nil")
		assert.Equal(t, updatedArticle.Title, article.Title, "Expected article with title %q", article.Title)
	})

	t.Run("Delete", func(t *testing.T) {
		err := repo.Delete("update")
		assert.NoError(t, err, "Expected no error")
		assert.Equal(t, len(repo.articles), 0, "Expected length of articles is zero")
	})
}
