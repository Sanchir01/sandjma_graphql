package response

import "github.com/Sanchir01/sandjma_graphql/internal/gql/model"

func NewInternalErrorProblem() model.InternalErrorProblem {
	return model.InternalErrorProblem{Message: "internal server error"}
}

func NewVersionMismatchProblem() model.VersionMismatchProblem {
	return model.VersionMismatchProblem{Message: "version mismatch"}
}
