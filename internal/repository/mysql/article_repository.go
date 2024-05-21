package mysql

import (
	"database/sql"
	model "go-clean-architecture/internal/domain"
	"time"
)

type ArticleMySqlRepository struct {
	DB *sql.DB
}

func NewArticleMySqlRepository(db *sql.DB) *ArticleMySqlRepository {
	return &ArticleMySqlRepository{DB: db}
}

func (repo *ArticleMySqlRepository) Create(article *model.Article) (*model.Article, error) {
	stmt, err := repo.DB.Prepare(`
		insert into article (title, body, userName, createdAt, updatedAt)
		values (?,?,?,?,?)
	`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		article.Title,
		article.Body,
		article.UserName,
		time.Now().Format("2006-01-02"),
		article.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (repo *ArticleMySqlRepository) GetById(title string) (*model.Article, error) {
	stmt, err := repo.DB.Prepare(`select title, body, userName, createdAt, updatedAt from article where title = ?`)
	if err != nil {
		return nil, err
	}
	var article model.Article
	rows, err := stmt.Query(title)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&article.Title, &article.Body, &article.UserName, &article.CreatedAt, &article.UpdatedAt)
	}

	return &article, nil
}

func (repo *ArticleMySqlRepository) GetAll() (res []*model.Article, err error) {
	stmt, err := repo.DB.Prepare(`select title, body, userName, createdAt, updatedAt from article`)
	if err != nil {
		return nil, err
	}
	var articles []*model.Article
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var article model.Article
		err = rows.Scan(&article.Title, &article.Body, &article.UserName, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}
	return articles, nil
}

func (repo *ArticleMySqlRepository) Update(title string, article *model.Article) (*model.Article, error) {
	article.UpdatedAt = time.Now().Format("2006-01-02")
	_, err := repo.DB.Exec("update article set body = ?, userName = ?, updatedAt = ? where title = ?", article.Body, article.UserName, article.UpdatedAt, article.Title)
	if err != nil {
		return nil, err
	}
	return article, err
}

func (repo *ArticleMySqlRepository) Delete(title string) error {
	_, err := repo.DB.Exec("delete from article where title = ?", title)
	if err != nil {
		return err
	}
	return nil
}
