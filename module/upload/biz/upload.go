package biz

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"path/filepath"
	"simple-rest-api/common"
	"simple-rest-api/component/uploadprovider"
	"simple-rest-api/module/upload/model"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	fileExt := filepath.Ext(fileName) // "img.jpg" => ".jpg"
	println(fileExt)
	w, h, err := getImageDimension(fileBytes, fileExt)

	if err != nil {
		return nil, model.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, model.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	//img.CloudName = "s3" // should be set in provider
	img.Extension = fileExt

	//if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	//	// delete img on S3
	//	return nil, uploadmodel.ErrCannotSaveFile(err)
	//}

	return img, nil
}

func getImageDimension(reader io.Reader, fileExt string) (int, int, error) {
	var img image.Config
	var err error
	if fileExt == ".png" {
		img, err = png.DecodeConfig(reader)
	} else if fileExt == ".jpeg" || fileExt == ".jpg" {
		img, err = jpeg.DecodeConfig(reader)
	}
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
