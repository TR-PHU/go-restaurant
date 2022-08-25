package restaurantbiz

import (
	"context"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	UpdateDataByCondition(ctx context.Context, conditions map[string]interface{}, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	err := biz.store.UpdateDataByCondition(ctx, map[string]interface{}{"id": id}, data)
	return err
}
