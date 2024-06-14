package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	runtime "github.com/Sanchir01/sandjma_graphql/internal/gql/generated"
)

// Query returns runtime.QueryResolver implementation.
func (r *Resolver) Query() runtime.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
