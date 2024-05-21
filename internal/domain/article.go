package domain

type Article struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserName  string `json:"userName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewArticle(title string, body string, userName string, createdAt string, updatedAt string) *Article {
	article := &Article{
		Title:     title,
		Body:      body,
		UserName:  userName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	return article
}
