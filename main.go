package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple-rest-api/component"
	"simple-rest-api/component/uploadprovider"
	"simple-rest-api/middleware"
	"simple-rest-api/module/restaurant/transport/ginrestaurant"
	"simple-rest-api/module/upload/transport/ginupload"
	userstorage "simple-rest-api/module/user/storage"
	"simple-rest-api/module/user/transport/ginuser"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, s3Provider uploadprovider.UploadProvider, secretKey string) error {
	userstore := userstorage.NewSQLStore(db)
	appCtx := component.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Set up static file
	//r.Static("/static", "./static")

	// CRUD
	restaurants := r.Group("/restaurant")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	uploads := r.Group("/upload")
	{
		uploads.POST("", ginupload.Upload(appCtx))
	}

	user := r.Group("/user")
	{
		user.POST("/register", ginuser.Register(appCtx))
		user.POST("/login", ginuser.Login(appCtx))
		user.GET("/profile", middleware.RequiredAuth(appCtx, userstore), ginuser.Profile(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
