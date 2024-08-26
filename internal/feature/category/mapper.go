package featureCategory

import "github.com/Sanchir01/sandjma_graphql/internal/gql/model"

func MapCategoryToGqlModel(categories []model.Category) (item []*model.Category, err error) {
	categoriesChan := make(chan *model.Category, len(categories))
	var categoryPtrs []*model.Category

	go func() {
		for i := range categories {
			categoriesChan <- &categories[i]
		}
		close(categoriesChan)
	}()

	for category := range categoriesChan {
		categoryPtrs = append(categoryPtrs, category)
	}
	return categoryPtrs, nil

}
