package somethingdataloader

import (
	"context"
	"fmt"

	dataloader "github.com/graph-gophers/dataloader/v7"
	"github.com/pkg/errors"
	"github.com/real-mielofon/dataloader-test/internal/model"
	"github.com/samber/lo"
)

func (d *Dataloader) GetSomething(ctx context.Context, id int64) (string, error) {

	resp, err := d.somethingDataloader.Load(ctx, id)()
	if err != nil {
		return "", fmt.Errorf("d.somethingDataloader.Load: %w", err)
	}

	return resp, nil
}

func NewSomethingDataloaderDataloader(client client) *dataloader.Loader[int64, string] {
	fn := func(ctx context.Context, ids []int64) []*dataloader.Result[string] {
		elements, err := client.GetSomething(ctx, ids)
		if err != nil {
			res := lo.Map(ids, func(_ int64, _ int) *dataloader.Result[string] {
				return &dataloader.Result[string]{
					Error: errors.Wrap(err, "client.GetSomething"),
				}
			})

			return res
		}

		elementsMap := lo.SliceToMap(elements, func(e model.Something) (int64, model.Something) {
			return e.ID, e
		})

		results := lo.Map(ids, func(id int64, _ int) *dataloader.Result[string] {
			element := elementsMap[id]

			return &dataloader.Result[string]{
				Data: element.Value,
			}
		})

		return results
	}

	cache := &dataloader.NoCache[int64, string]{}
	loader := dataloader.NewBatchedLoader(
		fn,
		dataloader.WithCache[int64, string](cache),
		dataloader.WithInputCapacity[int64, string](200),
	)

	return loader
}
