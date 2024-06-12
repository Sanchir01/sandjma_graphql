package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"

	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
)

// CreateProduct is the resolver for the createProduct field.
func (r *productMutationResolver) CreateProduct(ctx context.Context, obj *model.ProductMutation, input *model.CreateProductInput) (model.ProductCreateResult, error) {
	r.Logger.Info("CreateProduct", input)

	id, err := r.ProductStr.CreateProduct(ctx, input)
	if err != nil {
		return nil, err
	}
	return model.ProductCreateOk{Products: id}, nil
}
