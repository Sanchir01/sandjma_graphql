package directive

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	customMiddleware "github.com/Sanchir01/sandjma_graphql/internal/handlers/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type RoleDirectiveFunc = func(ctx context.Context, obj interface{}, next graphql.Resolver, role *model.Role) (res interface{}, err error)

func RoleDirective() RoleDirectiveFunc {
	return func(
		ctx context.Context,
		obj interface{},
		next graphql.Resolver,
		role *model.Role,
	) (res interface{}, err error) {

		ctxUserID, err := customMiddleware.GetJWTClaimsFromCtx(ctx)
		if err != nil {
			return nil, &gqlerror.Error{Message: "Unauthorized"}
		}

		if ctxUserID == nil {
			return nil, &gqlerror.Error{Message: "Unauthorized: user ID is nil"}
		}

		if role == nil {
			return nil, &gqlerror.Error{Message: "Role is nil"}
		}

		if ctxUserID.Role == *role {
			return next(ctx)
		}

		return nil, &gqlerror.Error{Message: "Role not admin"}
	}
}
