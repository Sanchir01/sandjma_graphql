package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"github.com/google/uuid"

	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *categoryMutationResolver) CreateCategory(ctx context.Context, obj *model.CategoryMutation, input *model.CreateCategoryInput) (model.CategoryCreateResult, error) {
	//newSlug, err := utils.Slugify(input.Name)
	//if err != nil {
	//	return response.NewInternalErrorProblem("error generating slug category"), err
	//}
	//id, err := r.CategoryStr.CreateCategory(ctx, input, newSlug)
	//if err != nil {
	//	return response.NewInternalErrorProblem("error creating category"), err
	//}
	return model.CategoryCreateOk{ID: uuid.New()}, nil
}
