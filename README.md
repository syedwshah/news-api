# GNews API

GNews API is a simple REST API that wraps the GNews API and returns JSON responses.

## Table of Contents

1.  [Getting Started](https://github.com/syedwshah/news-api#getting-started)
2.  [Endpoints]()
    1.  [/articles]()
    2.  [/articles/:size]()
    3.  [/articles/search/:query]()
    4.  [/articles/keyword/:keywords]()
3.  [Running Tests]()
4.  [Setting Up the API Token]()
5.  [Running the API Locally]()
6.  [Caching]()
7.  [Potential Improvements]()

## Getting Started

1.  Clone the repository: `git clone https://github.com/your_username/gnews-api`
2.  Change directory into the project folder: `cd gnews-api`
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

bashCopy code

`go test ./...`

This will run all tests in the project.

### Setting Up the API Token

To use the GNews API, you'll need an API token. You can sign up for a free token [here](https://gnews.io/signup). Once you have a token, create a `.env` file in the root directory of your project and add the following line, replacing `<your_token>` with your actual token:

makefileCopy code

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
