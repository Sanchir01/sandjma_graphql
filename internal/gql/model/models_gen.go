// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type GetAllProductResult interface {
	IsGetAllProductResult()
}

type ProblemInterface interface {
	IsProblemInterface()
	GetMessage() string
}

type ProductCreateResult interface {
	IsProductCreateResult()
}

type ProductInterface interface {
	IsProductInterface()
	GetID() string
	GetName() string
	GetPrice() string
	GetCratedAt() string
	GetUpdatedAt() string
	GetCategoryID() string
}

type VersionInterface interface {
	IsVersionInterface()
	GetVersion() uint
}

type CreateProductInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	CategoryID  int    `json:"categoryId"`
}

type GetAllProductsOk struct {
	Products []*Product `json:"products"`
}

func (GetAllProductsOk) IsGetAllProductResult() {}

type InternalErrorProblem struct {
	Message string `json:"message"`
}

func (InternalErrorProblem) IsProblemInterface()     {}
func (this InternalErrorProblem) GetMessage() string { return this.Message }

func (InternalErrorProblem) IsProductCreateResult() {}

func (InternalErrorProblem) IsGetAllProductResult() {}

type InvalidSortRankProblem struct {
	Message string `json:"message"`
}

func (InvalidSortRankProblem) IsProductCreateResult() {}

func (InvalidSortRankProblem) IsProblemInterface()     {}
func (this InvalidSortRankProblem) GetMessage() string { return this.Message }

type Mutation struct {
}

type Product struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CategoryID uuid.UUID `json:"category_id"`
	Version    uint      `json:"version"`
}

type ProductCreateOk struct {
	Products string `json:"products"`
}

func (ProductCreateOk) IsProductCreateResult() {}

type ProductMutation struct {
	CreateProduct ProductCreateResult `json:"createProduct"`
}

type ProductNotFoundProblem struct {
	Message string `json:"message"`
}

func (ProductNotFoundProblem) IsProblemInterface()     {}
func (this ProductNotFoundProblem) GetMessage() string { return this.Message }

func (ProductNotFoundProblem) IsProductCreateResult() {}

type ProductQuery struct {
	GetAllProduct GetAllProductResult `json:"getAllProduct"`
}

type Query struct {
}

type SortRankInput struct {
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type VersionMismatchProblem struct {
	Message string `json:"message"`
}

func (VersionMismatchProblem) IsProblemInterface()     {}
func (this VersionMismatchProblem) GetMessage() string { return this.Message }

type ArticleBlockFindSortEnum string

const (
	ArticleBlockFindSortEnumCreatedAtAsc  ArticleBlockFindSortEnum = "CREATED_AT_ASC"
	ArticleBlockFindSortEnumCreatedAtDesc ArticleBlockFindSortEnum = "CREATED_AT_DESC"
	ArticleBlockFindSortEnumSortRankAsc   ArticleBlockFindSortEnum = "SORT_RANK_ASC"
	ArticleBlockFindSortEnumSortRankDesc  ArticleBlockFindSortEnum = "SORT_RANK_DESC"
)

var AllArticleBlockFindSortEnum = []ArticleBlockFindSortEnum{
	ArticleBlockFindSortEnumCreatedAtAsc,
	ArticleBlockFindSortEnumCreatedAtDesc,
	ArticleBlockFindSortEnumSortRankAsc,
	ArticleBlockFindSortEnumSortRankDesc,
}

func (e ArticleBlockFindSortEnum) IsValid() bool {
	switch e {
	case ArticleBlockFindSortEnumCreatedAtAsc, ArticleBlockFindSortEnumCreatedAtDesc, ArticleBlockFindSortEnumSortRankAsc, ArticleBlockFindSortEnumSortRankDesc:
		return true
	}
	return false
}

func (e ArticleBlockFindSortEnum) String() string {
	return string(e)
}

func (e *ArticleBlockFindSortEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ArticleBlockFindSortEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ArticleBlockFindSortEnum", str)
	}
	return nil
}

func (e ArticleBlockFindSortEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
