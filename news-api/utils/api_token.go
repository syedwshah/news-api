package utils

import "os"

var APIToken string
var defaultToken string = "677d76e9d54a773265e916bc3346ffb8"

func init() {
    APIToken = os.Getenv("API_TOKEN")
    if APIToken == "" {
        APIToken = defaultToken
    }
}
