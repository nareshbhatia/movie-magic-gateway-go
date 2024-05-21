package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"log"

	"github.com/nareshbhatia/movie-magic-gateway-go/pkg/graph"
	"github.com/nareshbhatia/movie-magic-gateway-go/pkg/graph/model"
	moviepb "github.com/nareshbhatia/movie-magic-services-go/gen/go/movie/v1"
)

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context, input model.MoviesRequest) (*model.MoviesResponse, error) {
	req := &moviepb.ListMoviesRequest{}
	res, err := r.movieServiceClient.ListMovies(ctx, req)
	if err != nil {
		log.Printf("Error fetching movies: %v", err)
		return nil, err
	}

	// Map the gRPC response to the GraphQL model
	movies := make([]*model.Movie, len(res.Movies))
	for i, movie := range res.Movies {
		movies[i] = &model.Movie{
			ID:   movie.Id,
			Name: movie.Name,
			Cast: []*model.CastMember{
				{
					Person: &model.Person{
						ID:   "1",
						Name: "Actor 1",
					},
					Characters: []string{"Character 1", "Character 2"},
				},
			},
			Certificate: convertCertificate(movie.Certificate),
			Genres:      movie.Genres,
			Rank:        int(movie.Rank),
			RatingsSummary: &model.MovieRatingsSummary{
				AggregateRating: movie.RatingsSummary.AggregateRating,
				VoteCount:       int(movie.RatingsSummary.VoteCount),
			},
			ReleaseYear: int(movie.ReleaseYear),
		}
	}

	moviesResponse := &model.MoviesResponse{
		Movies: movies,
		PageInfo: &model.PaginationInfo{
			TotalPages:      1,
			TotalItems:      2,
			Page:            1,
			PerPage:         10,
			HasNextPage:     false,
			HasPreviousPage: false,
		},
	}

	return moviesResponse, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
