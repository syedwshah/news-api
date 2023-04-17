package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/news-api/models"
	"example.com/news-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// Endpoint for fetching news articles by size
func Articles(c *gin.Context) {
	apiKey := utils.APIToken
	fmt.Printf("API Key is %v\n",apiKey)

	client := resty.New()

	url := utils.GNewsAPIURL

	size := c.Param("size")

	// Set the API parameters
	params := map[string]string{
		"q":       "news",
		"lang":    "en",
		"country": "us",
		"max":     size,
		"token":   apiKey,
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

	// Return the articles as a JSON response
	c.JSON(http.StatusOK, response)
}