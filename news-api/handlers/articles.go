package handlers

import (
	"net/http"

	"example.com/news-api/models"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
    // Get query parameters from request
    q := c.Query("q")
    lang := c.Query("lang")
    country := c.Query("country")

    // Fetch articles from GNews API
    articles, err := models.GetArticles(q, lang, country)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
        return
    }

    // Return articles as JSON response
    c.JSON(http.StatusOK, articles)
}
