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

// Endpoint for searching by keywords
func SearchArticlesByKeyword(c *gin.Context) {
    apiKey := utils.APIToken

    keywords := c.Param("keywords")
    keywordArray := strings.Split(keywords, " ")

    client := resty.New()

    url := utils.GNewsAPIURL

    params := map[string]string{
        "q":     strings.Join(keywordArray, " OR "),
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

    // Filter articles by keyword
    var filteredArticles []models.Article
    for _, article := range response.Articles {
        for _, keyword := range keywordArray {
            if strings.Contains(strings.ToLower(article.Title), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(article.Description), strings.ToLower(keyword)) {
                filteredArticles = append(filteredArticles, article)
                break
            }
        }
    }

    // Return the filtered articles as a JSON response
    c.JSON(http.StatusOK, gin.H{"articles": filteredArticles})
}
