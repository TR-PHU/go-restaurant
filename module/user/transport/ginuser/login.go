package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/component/hasher"
	"simple-rest-api/component/tokenprovider/jwt"
	userbiz "simple-rest-api/module/user/biz"
	usermodel "simple-rest-api/module/user/model"
	userstorage "simple-rest-api/module/user/storage"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		store := userstorage.NewSQLStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)

		account, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
