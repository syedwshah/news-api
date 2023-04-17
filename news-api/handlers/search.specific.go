package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"example.com/news-api/models"
	"example.com/news-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// Endpoint for finding a news article by title or author
func SearchSpecificArticles(c *gin.Context) {
    apiKey := utils.APIToken    
    query := c.Param("query")

    client := resty.New()

    url := utils.GNewsAPIURL

    params := map[string]string{
        "q":     query,
        "lang":  "en",
        "token": apiKey,
    }

    // Send the API request
    resp, err := client.R().
        SetQueryParams(params).
        SetHeader("Accept", "application/json").
        Get(url)

    if err != nil {
        fmt.Println("Error:", err)
        c.JSON(500, gin.H{"error": "Internal server error"})
        return
    }
    if resp.StatusCode() != 200 {
        fmt.Println("Error: API returned non-200 status code:", resp.StatusCode())
        c.JSON(resp.StatusCode(), gin.H{"error": "Failed to fetch articles"})
        return
    }

    // Parse the API response body
    var response models.Response
    if err := json.Unmarshal(resp.Body(), &response); err != nil {
        fmt.Println("Error:", err)
        c.JSON(500, gin.H{"error": "Internal server error"})
        return
    }

    // Filter articles by title or author
    var filteredArticles []models.Article
    for _, article := range response.Articles {
        if strings.Contains(strings.ToLower(article.Title), strings.ToLower(query)) ||
            strings.Contains(strings.ToLower(article.Source.Name), strings.ToLower(query)) {
            filteredArticles = append(filteredArticles, article)
        }
    }

    // Return the filtered articles as a JSON response
    c.JSON(http.StatusOK, gin.H{"articles": filteredArticles})

}