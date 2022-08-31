package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/module/restaurant/biz"
	"simple-rest-api/module/restaurant/storage"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := biz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
