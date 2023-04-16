package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Url         string `json:"url"`
    ImageUrl    string `json:"image"`
    PublishedAt string `json:"publishedAt"`
    Source      struct {
        Name string `json:"name"`
    } `json:"source"`
}

func GetArticles(q string, lang string, country string) ([]Article, error) {
    // Make HTTP request to GNews API
    resp, err := http.Get(fmt.Sprintf("https://gnews.io/api/v4/search?q=%s&lang=%s&country=%s&token=YOUR_API_TOKEN", q, lang, country))

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    // Parse response JSON
    var result struct {
        Articles []Article `json:"articles"`
    }

    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result.Articles, nil
}
