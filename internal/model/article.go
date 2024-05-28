package model

type Article struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserName  string `json:"userName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ArticleUseCase interface {
	Create(article *Article) (*Article, error)
	GetById(title string) (res *Article, err error)
	GetAll() (res []*Article, err error)
	Update(title string, article *Article) (*Article, error)
	Delete(title string) error
}

type ArticleRepository interface {
	Create(article *Article) (*Article, error)
	GetById(title string) (*Article, error)
	GetAll() (res []*Article, err error)
	Update(title string, article *Article) (*Article, error)
	Delete(title string) error
}
