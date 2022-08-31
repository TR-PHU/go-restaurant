package ginupload

import (
	"github.com/gin-gonic/gin"
	"simple-rest-api/common"
	"simple-rest-api/component"
	"simple-rest-api/module/upload/biz"
)

func Upload(appCtx component.AppContext) func(context2 *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// folder storage in provider
		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // if err -> close file

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//imgStore := uploadstorage.NewSQLStore(db)
		biz := biz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))

		//if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
		//	panic(err)
		//}
		//
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
		//	Id:        0,
		//	Url:       "http://localhost:8080/static/" + fileHeader.Filename,
		//	Width:     0,
		//	Height:    0,
		//	CloudName: "local",
		//	Extension: "img",
		//}))
	}
}
