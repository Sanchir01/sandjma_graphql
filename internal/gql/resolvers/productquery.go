package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	runtime "github.com/Sanchir01/sandjma_graphql/internal/gql/generated"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
)

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) (*model.ProductQuery, error) {
	return &model.ProductQuery{}, nil
}

// ProductQuery returns runtime.ProductQueryResolver implementation.
func (r *Resolver) ProductQuery() runtime.ProductQueryResolver { return &productQueryResolver{r} }

type productQueryResolver struct{ *Resolver }