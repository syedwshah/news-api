# GNews API

GNews API is a simple REST API that wraps the GNews API and returns JSON responses.

## Table of Contents

1.  [Getting Started](https://github.com/syedwshah/news-api#getting-started)
2.  [Endpoints](https://github.com/syedwshah/news-api#endpoints)
    1.  [/articles](https://github.com/syedwshah/news-api#articles)
    2.  [/articles/:size](https://github.com/syedwshah/news-api#articlessize)
    3.  [/articles/search/:query](https://github.com/syedwshah/news-api#articlessearchquery)
    4.  [/articles/keyword/:keywords](https://github.com/syedwshah/news-api#articleskeywordkeywords)
3.  [Running Tests](https://github.com/syedwshah/news-api#running-tests)
4.  [Setting Up the API Token](https://github.com/syedwshah/news-api#setting-up-the-api-token)
5.  [Running the API Locally](https://github.com/syedwshah/news-api#running-the-api-locally)
6.  [Caching](https://github.com/syedwshah/news-api#caching)
7.  [Potential Improvements](https://github.com/syedwshah/news-api#potential-improvements)
8.  [Manual Testing](https://github.com/syedwshah/news-api#manual-testing)

## Getting Started

1.  Clone the repository: `git clone https://github.com/your_username/gnews-api`
2.  Change directory into the project folder: `cd news-api`
3.  Install dependencies: `go install`
4.  Create a .env file with the following contents: `API_KEY=<your_token>`, replacing `<your_token>` with your GNews API token. If you do not have a token, the default token in `handlers/handlers.go` will be used.
5.  Start the server: `go run main.go`

## Endpoints

### /articles

- **Description**: Returns the top 10 headlines from GNews API by default.
- **Method**: GET
- **URL**: /articles
- **Query Parameters**:
  - `topic`: Optional. Filter the articles by topic. For example: `/articles?topic=world`
- **Success Response**:
  - Code: 200
  - Content: JSON response containing articles

### /articles/:size

- **Description**: Returns the top `size` headlines from GNews API.
- **Method**: GET
- **URL**: /articles/:size
- **URL Parameters**:
  - `size`: Required. Number of articles to return. For example: `/articles/5`
- **Query Parameters**:
  - `topic`: Optional. Filter the articles by topic. For example: `/articles/5?topic=world`
- **Success Response**:
  - Code: 200
  - Content: JSON response containing articles

### /articles/search/:query

- **Description**: Returns articles containing the specified search query.
- **Method**: GET
- **URL**: /articles/search/:query
- **URL Parameters**:
  - `query`: Required. The search query to match against article titles.
- **Success Response**:
  - Code: 200
  - Content: JSON response containing articles matching the search query.

### /articles/keyword/:keywords

- **Description**: Returns articles containing the specified keywords.
- **Method**: GET
- **URL**: /articles/keyword/:keywords
- **URL Parameters**:
  - `keywords`: Required. The keywords to match against article descriptions.
- **Query Parameters**:
  - `topic`: Optional. Filter the articles by topic. For example: `/articles/keyword/tech?topic=world`
- **Success Response**:
  - Code: 200
  - Content: JSON response containing articles containing the specified keywords.

## Running Tests

To run the tests, navigate to the root directory of the project and use the following command:

`go test ./...`

This will run all tests in the project.

### Setting Up the API Token

To use the GNews API, you'll need an API token. You can sign up for a free token [here](https://gnews.io/signup). Once you have a token, create a `.env` file in the root directory of your project and add the following line, replacing `<your_token>` with your actual token:

`API_KEY=<your_token>`

### Running the API Locally

To run the API locally, follow these steps:

1.  Clone the repository: `git clone https://github.com/your_username/gnews-api`
2.  Change directory into the project folder: `cd gnews-api`
3.  Install dependencies: `npm install`
4.  Start the server: `npm start`

### Caching

The API uses caching to improve performance. The first time an endpoint is hit, the response is cached. Subsequent requests to the same endpoint will return the cached response until the cache expires.

The expiration time for the cache is set to 1 hour by default. You can modify this value by changing the `CACHE_EXPIRATION_TIME` constant in your code.

If you want to clear the cache manually, you can do so by sending a GET request to the `/cache/clear` endpoint.

### Potential Improvements

Here are some potential improvements that could be made to the API:

- Adding pagination to the `/articles` endpoint to allow for retrieval of more than 10 articles at a time
- Adding additional query parameters to the `/articles` endpoint, such as sorting by date or relevance
- Implementing user authentication to allow for personalized news feeds and saved articles
- Adding support for other news APIs to expand the sources of news articles available
- Optimizing the caching mechanism to prevent cache stampedes and reduce the amount of memory used by the cache

## Manual Testing

Make sure the server is running using "go run main.go" in the terminal. To manually test the API endpoints, you can use `curl` or any other HTTP client tool. Here are some examples:

### Get articles:

`curl http://localhost:8080/articles  # default to 10 curl http://localhost:8080/articles/5 # get 5 articles`

### Get specific articles

#### Get by title:

`curl http://localhost:8080/articles/search/"How%20to%20download%20and%20add%20EPUB%20books%20to%20your%20Amazon%20Kindle"`

#### Get by name of source:

`curl -X GET "http://localhost:8080/articles/search/android%20central"`

#### Get by description keywords:

`curl -X GET "http://localhost:8080/articles/keyword/how%20to%20download%20and%20add"`

You can use these examples as a starting point and modify them according to your specific testing needs.
