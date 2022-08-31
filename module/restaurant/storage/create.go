package storage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/module/restaurant/model"
)

// This method means struct sqlStore implement method Create of interface CreateRestaurantStore
func (s *sqlStore) Create(ctx context.Context, data *model.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
