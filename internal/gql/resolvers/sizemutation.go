package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	runtime "github.com/Sanchir01/sandjma_graphql/internal/gql/generated"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
)

// Size is the resolver for the size field.
func (r *mutationResolver) Size(ctx context.Context) (*model.SizeMutation, error) {
	return &model.SizeMutation{}, nil
}

// SizeMutation returns runtime.SizeMutationResolver implementation.
func (r *Resolver) SizeMutation() runtime.SizeMutationResolver { return &sizeMutationResolver{r} }

type sizeMutationResolver struct{ *Resolver }
