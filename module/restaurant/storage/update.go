package storage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/module/restaurant/model"
)

// This method means struct sqlStore implement method UpdateData of interface UpdateRestaurantStore
func (s *sqlStore) UpdateData(ctx context.Context, id int, data *model.RestaurantUpdate) error {

	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
