package main

import (
	"database/sql"
	"fmt"
	"go-clean-architecture/internal/repository/mysql"
	"net/http"
	"os"

	_articleController "go-clean-architecture/internal/controller"
	_articleUseCase "go-clean-architecture/internal/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var (
	dbUser = os.Getenv("MYSQL_USER")
	dbPass = os.Getenv("MYSQL_PASSWORD")
	dbName = os.Getenv("MYSQL_DATABASE")
	dbHost = os.Getenv("MYSQL_HOST")
	dbPort = os.Getenv("MYSQL_PORT")
)

func main() {
	router := mux.NewRouter()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset-utf8mb4&parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to MySQL:", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping MySQL:", err)
		os.Exit(1)
	}

	defer db.Close()

	articleRepository := mysql.NewArticleMySqlRepository(db)
	articleUseCase := _articleUseCase.NewArticleUseCase(articleRepository)
	_articleController.NewArticleController(router, articleUseCase)

	fmt.Println("listening localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
