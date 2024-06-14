package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"log/slog"

	featureCategory "github.com/Sanchir01/sandjma_graphql/internal/feature/category"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/Sanchir01/sandjma_graphql/pkg/lib/api/response"
)

// GetAllCategory is the resolver for the getAllCategory field.
func (r *categoryQueryResolver) GetAllCategory(ctx context.Context, obj *model.CategoryQuery) (model.CategoryGetAllResult, error) {
	categoryStr, err := r.Resolver.CategoryStr.GetAllCategory(ctx)
	if err != nil {
		r.Logger.Error("GetAllCategory error", slog.String("error", err.Error()))
		return response.NewInternalErrorProblem("error for get all category db"), nil
	}

	categoriesGql, err := featureCategory.MapCategoryToGqlModel(categoryStr)

	if err != nil {
		r.Logger.Error("GetAllCategory mapping gql model error", slog.String("error", err.Error()))
		return response.NewInternalErrorProblem("error for mapping category gql"), nil
	}

	return model.CategoryGetAllOk{Category: categoriesGql}, nil
}
