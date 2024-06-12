package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"

	runtime "github.com/Sanchir01/sandjma_graphql/internal/gql/generated"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
)

// Category is the resolver for the category field.
func (r *mutationResolver) Category(ctx context.Context) (*model.CategoryMutation, error) {
	return &model.CategoryMutation{}, nil
}

// CategoryMutation returns runtime.CategoryMutationResolver implementation.
func (r *Resolver) CategoryMutation() runtime.CategoryMutationResolver {
	return &categoryMutationResolver{r}
}

type categoryMutationResolver struct{ *Resolver }
