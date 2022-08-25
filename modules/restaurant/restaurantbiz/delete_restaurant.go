package restaurantbiz

import "context"

type DeleteRestaurantStore interface {
	DeleteDataByCondition(ctx context.Context, conditions map[string]interface{}) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	if err := biz.store.DeleteDataByCondition(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
