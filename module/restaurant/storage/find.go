package storage

import (
	"context"
	"gorm.io/gorm"
	"simple-rest-api/common"
	"simple-rest-api/module/restaurant/model"
)

// This method means struct sqlStore implement method FindDataByCondition of interface GetRestaurantStore
func (s *sqlStore) FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.Restaurant, error) {
	var result model.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &result, nil
}
