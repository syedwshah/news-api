package main

import (
	"log"

	"example.com/news-api/handlers"
	"example.com/news-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitCache()

	
	r := gin.Default()

	r.GET("/articles", handlers.Articles)
	r.GET("/articles/:size", handlers.Articles)
	r.GET("/articles/search/:query", handlers.SearchSpecificArticles)
    r.GET("/articles/keyword/:keywords", handlers.SearchArticlesByKeyword)

    r.Run("localhost:8080")
}
