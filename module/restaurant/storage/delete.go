package storage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/module/restaurant/model"
)

// This method means struct sqlStore implement method SoftDeleteData of interface SoftDeleteRestaurant
func (s *sqlStore) SoftDeleteData(ctx context.Context, id int) error {

	db := s.db

	if err := db.Table(model.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
