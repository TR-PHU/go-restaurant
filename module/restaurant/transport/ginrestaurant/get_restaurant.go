package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/module/restaurant/biz"
	"simple-rest-api/module/restaurant/storage"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := biz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
