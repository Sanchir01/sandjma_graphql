package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	featureSize "github.com/Sanchir01/sandjma_graphql/internal/feature/size"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/Sanchir01/sandjma_graphql/pkg/lib/api/response"
)

// GetAllSizes is the resolver for the getAllSizes field.
func (r *sizeQueryResolver) GetAllSizes(ctx context.Context, obj *model.SizeQuery) (model.GetAllSizeResult, error) {
	sizes, err := r.SizeStr.GetAllSizes(ctx)
	if err != nil {
		return response.NewInternalErrorProblem(""), err
	}
	mappingSuzes, err := featureSize.MapManySizesToGqlModels(sizes)
	r.Logger.Info("size", mappingSuzes)
	return model.GetAllSizeOk{Sizes: mappingSuzes}, nil
}
