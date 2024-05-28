package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	handler "go-clean-architecture/internal/handler"
	"go-clean-architecture/internal/repository/mysql"
	usecase "go-clean-architecture/internal/usecase"

	httpSwagger "github.com/swaggo/http-swagger"

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

// @title go-clean-architecture
// @version 0.0.1
// @description go-clean-architecture
// @host localhost:8080
// @basePath /api
func main() {
	router := mux.NewRouter()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to MySQL:", err)
		os.Exit(1)
	}

	defer db.Close()

	handler.NewArticleHandler(router, usecase.NewArticleService(mysql.NewArticleMySqlRepository(db)))

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	fmt.Println("listening localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
