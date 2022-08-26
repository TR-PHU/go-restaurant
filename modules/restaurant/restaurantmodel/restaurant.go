package restaurantmodel

import (
	"errors"
	"simple-rest-api/common"
	"strings"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Address         string `json:"address" gorm:"column:address;"`
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:address;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"id" gorm:"column:id;"`
	Name            string `json:"name" gorm:"column:name;"`
	Address         string `json:"address" gorm:"column:address;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if len(data.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
