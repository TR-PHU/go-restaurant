package component

import (
	"gorm.io/gorm"
	"simple-rest-api/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider}
}

//This 2 methods means type appCtx implement the interface AppContext
func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}
