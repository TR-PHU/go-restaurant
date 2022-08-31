package biz

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *model.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *model.RestaurantCreate) error {

	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
