package restaurantstore

import (
	"context"
	"simple-rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DeleteDataByCondition(ctx context.Context, conditions map[string]interface{}) error {

	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
