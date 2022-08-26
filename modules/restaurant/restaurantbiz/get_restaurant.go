package restaurantbiz

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, err)
	}
	return data, nil
}
