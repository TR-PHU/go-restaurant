package userbiz

import (
	"context"
	"simple-rest-api/common"
	usermodel "simple-rest-api/module/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBiz(registerStorage RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		//if user.Status == 0 {
		//	return error user has been disable
		//}

		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
