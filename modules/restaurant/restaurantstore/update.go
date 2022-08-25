package restaurantstore

import (
	"context"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateDataByCondition(ctx context.Context, conditions map[string]interface{}, data *restaurantmodel.RestaurantUpdate) error {

	db := s.db

	if err := db.Table(restaurantmodel.RestaurantUpdate{}.TableName()).Where(conditions).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
