package somethingdataloader

import (
	"context"
	"time"

	dataloader "github.com/graph-gophers/dataloader/v7"
	"github.com/real-mielofon/dataloader-test/internal/model"
)

type client interface {
	GetSomething(ctx context.Context, ids []int64) ([]model.Something, error)
}

type Dataloader struct {
	somethingDataloader *dataloader.Loader[int64, string]
}

// New создает новый экземпляр Dataloader
func New(client client) *Dataloader {
	return newWithTimeout(client, 250)
}

func newWithTimeout(client client, wait time.Duration) *Dataloader {
	return &Dataloader{
		somethingDataloader: NewSomethingDataloaderDataloader(client),
	}
}
