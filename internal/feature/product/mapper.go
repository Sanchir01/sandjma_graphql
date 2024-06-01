package fetureProduct

import "github.com/Sanchir01/sandjma_graphql/internal/gql/model"

func MapManyProductsToGqlModels(products []model.Product) (items []*model.Product, err error) {
	productChan := make(chan *model.Product, len(products))
	var productPtrs []*model.Product

	go func() {
		for i := range products {
			productChan <- &products[i]
		}
		close(productChan)
	}()

	// Читаем из канала и заполняем срез productPtrs
	for product := range productChan {
		productPtrs = append(productPtrs, product)
	}
	return productPtrs, nil
}
