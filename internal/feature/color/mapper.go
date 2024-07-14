package featureColor

import "github.com/Sanchir01/sandjma_graphql/internal/gql/model"

func MapManyColorsToGqlModels(colors []model.Color) ([]*model.Color, error) {
	colorsChan := make(chan *model.Color, len(colors))
	var productPtrs []*model.Color

	go func() {
		for i := range colors {
			colorsChan <- &colors[i]
		}
		close(colorsChan)
	}()

	for product := range colorsChan {
		productPtrs = append(productPtrs, product)
	}
	return productPtrs, nil
}
