package test

import (
	_article "go-clean-architecture/internal/domain"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewArticle(t *testing.T) {
	Convey("New Article", t, func() {
		article := _article.NewArticle("title", "body", "username", time.Now().Format("2006-01-02 15:04:05"), time.Now().Add(5*time.Minute).Format("2006-01-02 15:04:05"))

		Convey("Article Validation", func() {
			So(article.Title, ShouldEqual, "title")
		})
	})
}
