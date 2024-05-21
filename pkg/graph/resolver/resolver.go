package resolver

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	moviepb "github.com/nareshbhatia/movie-magic-services-go/gen/go/movie/v1"
)

type Resolver struct {
	movieServiceClient moviepb.MovieServiceClient
}

func NewResolver(movieServiceClient moviepb.MovieServiceClient) *Resolver {
	return &Resolver{movieServiceClient: movieServiceClient}
}
