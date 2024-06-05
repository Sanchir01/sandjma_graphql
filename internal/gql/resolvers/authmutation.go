package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	runtime "github.com/Sanchir01/sandjma_graphql/internal/gql/generated"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
)

// Auth is the resolver for the user field.
func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthMutation, error) {

	return &model.AuthMutation{}, nil
}

// AuthMutation returns runtime.AuthMutationResolver implementation.
func (r *Resolver) AuthMutation() runtime.AuthMutationResolver { return &authMutationResolver{r} }

type authMutationResolver struct{ *Resolver }