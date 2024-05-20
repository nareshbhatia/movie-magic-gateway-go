package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"

	"github.com/nareshbhatia/movie-magic-gateway-go/pkg/graph"
	"github.com/nareshbhatia/movie-magic-gateway-go/pkg/graph/model"
)

// Movies is the resolver for the movies field.
func (r *queryResolver) Movies(ctx context.Context, filter model.MoviesRequest) (*model.MoviesResponse, error) {
	panic(fmt.Errorf("not implemented: Movies - movies"))
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
