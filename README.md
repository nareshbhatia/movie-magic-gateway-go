# Movie Magic GraphQL Gateway in Go

## Getting started

1. Make sure you have Go installed.
2. In your shell, run `go run server.go` to start the server.
3. Open GraphiQL at http://localhost:8080
4. Execute the following query:

```graphql
# ----- query -----
query ListMovies($input: MoviesRequest!) {
  movies(input: $input) {
    movies {
      id
      name
      cast {
        person {
          id
          name
        }
        characters
      }
      certificate
      genres
      rank
      releaseYear
      ratingsSummary {
        aggregateRating
      }
    }
    pageInfo {
      totalPages
      totalItems
      page
      perPage
      hasNextPage
      hasPreviousPage
    }
  }
}

# ----- variables -----
{
  "input": {
    "sort": "SORT_PARAM_RANK_ASC",
    "pageSpec": {
      "page": 1,
      "perPage": 10
    }
  }
}
```

You should see top 10 movies in the result.

## Running gqlgen

```shell
go run github.com/99designs/gqlgen generate
```
