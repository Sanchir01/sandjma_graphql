package featureSize

import "github.com/Sanchir01/sandjma_graphql/internal/gql/model"

func MapManySizesToGqlModels(sizes []model.Size) ([]*model.Size, error) {
	// Создаем канал для передачи указателей *Size
	sizesChan := make(chan *model.Size, len(sizes))
	var sizesPtrs []*model.Size

	go func() {

		for i := range sizes {

			sizesChan <- &sizes[i]
		}
		close(sizesChan)
	}()

	for size := range sizesChan {
		sizesPtrs = append(sizesPtrs, size)
	}

	return sizesPtrs, nil
}
