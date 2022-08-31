package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/module/restaurant/biz"
	"simple-rest-api/module/restaurant/model"
	"simple-rest-api/module/restaurant/storage"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(err)
		}

		var data model.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := biz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
