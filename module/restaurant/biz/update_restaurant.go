package biz

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*model.Restaurant, error)
	UpdateData(ctx context.Context, id int, data *model.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *model.RestaurantUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
