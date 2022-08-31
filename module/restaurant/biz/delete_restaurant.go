package biz

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.Restaurant, error)
	SoftDeleteData(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {

	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(model.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, nil)
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, nil)
	}

	return nil
}
