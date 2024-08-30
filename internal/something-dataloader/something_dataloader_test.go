package somethingdataloader

import (
	"context"
	"testing"
	"time"

	"github.com/real-mielofon/dataloader-test/internal/model"
	"github.com/real-mielofon/dataloader-test/internal/something-dataloader/_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDataloader_GetSomething(t *testing.T) {
	ctrl := gomock.NewController(t)
	regClientMock := _mocks.NewMockclient(ctrl)
	d := newWithTimeout(regClientMock, 10*time.Millisecond)

	ctx := context.Background()
	elementID := int64(2)

	elements := []model.Something{
		{
			ID:    elementID,
			Value: "test1",
		},
	}

	regClientMock.EXPECT().GetSomething(ctx, []int64{elementID}).Return(elements, nil)

	got, err := d.GetSomething(ctx, elementID)

	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, elements[0].Value, got)
}

func TestDataloader_GetSomething_ErrorInExcept(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	regClientMock := _mocks.NewMockclient(ctrl)
	d := newWithTimeout(regClientMock, 10*time.Millisecond)

	ctx := context.Background()
	elementID := int64(2)
	wrongElementID := int64(200)

	elements := []model.Something{
		{
			ID:    wrongElementID,
			Value: "test1",
		},
	}

	// error in EXPECT - elementID + 10
	regClientMock.EXPECT().GetSomething(ctx, []int64{wrongElementID}).Return(elements, nil)

	got, err := d.GetSomething(ctx, elementID)

	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, elements[0].Value, got)
}
