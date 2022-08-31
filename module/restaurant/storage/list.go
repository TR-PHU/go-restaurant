package storage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/module/restaurant/model"
)

// This method means struct sqlStore implement method ListDataByCondition of interface ListRestaurantStore
func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Restaurant, error) {
	var result []model.Restaurant
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(model.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Order("id desc").
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
